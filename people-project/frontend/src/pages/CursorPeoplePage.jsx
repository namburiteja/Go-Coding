import { useCallback, useEffect, useState } from "react";
import "../App.css";

import {
  getPeopleCursor,
  getPersonById,
  searchPeopleByName,
} from "../api/peopleApi";

import PeopleTable from "../components/PeopleTable";

function CursorPeoplePage({ onPersonClick }) {
  const [people, setPeople] = useState([]);

  const [nextCursor, setNextCursor] = useState("");
  const [hasNext, setHasNext] = useState(true);

  const [loading, setLoading] = useState(false);
  const [loadingMore, setLoadingMore] = useState(false);

  const [searchText, setSearchText] = useState("");
  const [searchResults, setSearchResults] = useState([]);

  const [searchLoading, setSearchLoading] = useState(false);

  const [error, setError] = useState("");

  const limit = 10;

  /*
   * ==========================================
   * INITIAL CURSOR LOAD
   * ==========================================
   */

  const loadInitialPeople = useCallback(async () => {
    try {
      setLoading(true);
      setError("");

      const data = await getPeopleCursor("", limit);

      setPeople(data.data);
      setNextCursor(data.nextCursor || "");
      setHasNext(data.hasNext);
    } catch (err) {
      console.error(err);
      setError("Failed to load people");
    } finally {
      setLoading(false);
    }
  }, []);

  useEffect(() => {
    loadInitialPeople();
  }, [loadInitialPeople]);

  /*
   * ==========================================
   * LOAD MORE
   * ==========================================
   */

  const loadMorePeople = useCallback(async () => {
    if (!hasNext || loadingMore || !nextCursor) {
      return;
    }

    try {
      setLoadingMore(true);
      setError("");

      const data = await getPeopleCursor(
        nextCursor,
        limit
      );

      setPeople((currentPeople) => [
        ...currentPeople,
        ...data.data,
      ]);

      setNextCursor(data.nextCursor || "");
      setHasNext(data.hasNext);
    } catch (err) {
      console.error(err);
      setError("Failed to load more people");
    } finally {
      setLoadingMore(false);
    }
  }, [hasNext, loadingMore, nextCursor]);

  /*
   * ==========================================
   * LIVE SEARCH
   * ==========================================
   */

  useEffect(() => {
    const value = searchText.trim();

    if (value === "") {
      setSearchResults([]);
      setSearchLoading(false);
      setError("");
      return;
    }

    if (value.length < 2) {
      setSearchResults([]);
      setSearchLoading(false);
      return;
    }

    const timer = setTimeout(async () => {
      try {
        setSearchLoading(true);
        setError("");

        const looksLikePlayerID =
          /^[a-zA-Z0-9]+$/.test(value) &&
          /\d/.test(value) &&
          value.length >= 6;

        if (looksLikePlayerID) {
          try {
            const person = await getPersonById(value);

            setSearchResults([person]);
            return;
          } catch (err) {
            if (err.response?.status !== 404) {
              throw err;
            }
          }
        }

        const results = await searchPeopleByName(value);

        setSearchResults(results);
      } catch (err) {
        console.error(err);

        setSearchResults([]);
        setError("Failed to search people");
      } finally {
        setSearchLoading(false);
      }
    }, 300);

    return () => clearTimeout(timer);
  }, [searchText]);

  /*
   * ==========================================
   * SEARCH
   * ==========================================
   */

  const handleSearchChange = (event) => {
    setSearchText(event.target.value);
    setError("");
  };

  const handleClearSearch = () => {
    setSearchText("");
    setSearchResults([]);
    setError("");
  };

  /*
   * ==========================================
   * DATA TO DISPLAY
   * ==========================================
   */

  const isSearching =
    searchText.trim().length >= 2;

  const tableData = isSearching
    ? searchResults
    : people;

  /*
   * ==========================================
   * UI
   * ==========================================
   */

  return (
    <div className="app">
      <div className="container">

        <div className="header">
          <div>
            <h1>People API</h1>

            <p>
              Cursor Pagination • Infinite Scroll
            </p>
          </div>
        </div>

        <div className="search-wrapper">
          <div className="search-bar">

            <input
              type="text"
              placeholder="Search by Player ID or Name..."
              value={searchText}
              onChange={handleSearchChange}
            />

            {searchText && (
              <button
                type="button"
                className="search-clear"
                onClick={handleClearSearch}
                aria-label="Clear search"
              >
            
              </button>
            )}

          </div>
        </div>

        {error && (
          <div className="error">
            {error}
          </div>
        )}

        {loading || searchLoading ? (
          <div className="loading">
            {searchLoading
              ? "Searching..."
              : "Loading..."}
          </div>
        ) : (
          <PeopleTable
            people={tableData}
            onPersonClick={onPersonClick}
            infiniteScroll={!isSearching}
            hasNext={hasNext}
            loadingMore={loadingMore}
            onLoadMore={loadMorePeople}
          />
        )}

      </div>
    </div>
  );
}

export default CursorPeoplePage;