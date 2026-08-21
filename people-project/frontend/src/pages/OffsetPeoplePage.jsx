import { useEffect, useState } from "react";
import "../App.css";

import {
  getPeople,
  getPersonById,
  searchPeopleByName,
} from "../api/peopleApi";

import PeopleTable from "../components/PeopleTable";
import Pagination from "../components/Pagination";
import SortControls from "../components/SortControls";

function OffsetPeoplePage({ onPersonClick }) {
  const [people, setPeople] = useState([]);

  const [page, setPage] = useState(1);
  const [limit] = useState(10);
  const [totalPages, setTotalPages] = useState(1);

  const [searchText, setSearchText] = useState("");
  const [searchResults, setSearchResults] = useState([]);

  const [loading, setLoading] = useState(false);
  const [searchLoading, setSearchLoading] = useState(false);

  const [error, setError] = useState("");

  const [sortBy, setSortBy] = useState("");
  const [sortOrder, setSortOrder] = useState("asc");

  /*
   * ==========================================
   * LOAD OFFSET PAGINATED PEOPLE
   * ==========================================
   */

  useEffect(() => {
    if (searchText.trim() !== "") {
      return;
    }

    loadPeople();
  }, [page, searchText, sortBy, sortOrder]);

  const loadPeople = async () => {
    try {
      setLoading(true);
      setError("");

      const data = await getPeople(page,limit,sortBy,sortOrder);

      setPeople(data.data);
      setTotalPages(data.totalPages);
    } catch (err) {
      console.error(err);
      setError("Failed to load people");
    } finally {
      setLoading(false);
    }
  };

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
    const value = event.target.value;

    setSearchText(value);
    setPage(1);
    setError("");
  };

  const handleClearSearch = () => {
    setSearchText("");
    setSearchResults([]);
    setError("");
    setPage(1);
  };

  /*
   * ==========================================
   * DATA TO DISPLAY
   * ==========================================
   */

  const tableData =
    searchText.trim().length >= 2
      ? searchResults
      : people;

  const isSearching =
    searchText.trim().length >= 2;

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
              Offset Pagination
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
                ×
              </button>
            )}

          </div>
          <SortControls
              sortBy={sortBy}
              sortOrder={sortOrder}
              onSortByChange={(value) => {
                setSortBy(value);
                setPage(1);
              }}
              onSortOrderChange={(value) => {
                setSortOrder(value);
                setPage(1);
              }}
            />
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
          <>
            <PeopleTable
              people={tableData}
              onPersonClick={onPersonClick}
            />

            {!isSearching && (
              <Pagination
                page={page}
                totalPages={totalPages}
                onPrevious={() =>
                  setPage((p) => p - 1)
                }
                onNext={() =>
                  setPage((p) => p + 1)
                }
              />
            )}
          </>
        )}

      </div>
    </div>
  );
}

export default OffsetPeoplePage;