import { useEffect, useState } from "react";
import "./App.css";

import {
  getPeople,
  getPersonById,
  searchPeopleByName,
} from "./api/peopleApi";

import PeopleTable from "./components/PeopleTable";
import Pagination from "./components/Pagination";
import PersonDetails from "./components/PersonDetails";

function App() {
  const [people, setPeople] = useState([]);

  const [page, setPage] = useState(1);
  const [limit] = useState(10);
  const [totalPages, setTotalPages] = useState(1);

  const [selectedPerson, setSelectedPerson] = useState(null);

  // One search input for both ID and name
  const [searchText, setSearchText] = useState("");

  // Search result data
  const [searchResults, setSearchResults] = useState([]);

  const [loading, setLoading] = useState(false);
  const [searchLoading, setSearchLoading] = useState(false);

  const [error, setError] = useState("");

  /*
   * ==========================================
   * LOAD NORMAL PAGINATED PEOPLE
   * ==========================================
   */

  useEffect(() => {
    // Don't load normal pagination while searching
    if (searchText.trim() !== "") {
      return;
    }

    loadPeople();
  }, [page, searchText]);

  const loadPeople = async () => {
    try {
      setLoading(true);
      setError("");

      const data = await getPeople(page, limit);

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

    /*
     * Empty search
     *
     * Return to normal paginated data.
     */
    if (value === "") {
      setSearchResults([]);
      setSearchLoading(false);
      setError("");
      return;
    }

    /*
     * Don't search until 2 characters
     */
    if (value.length < 2) {
      setSearchResults([]);
      setSearchLoading(false);
      return;
    }

    /*
     * Wait 300ms after user stops typing.
     *
     * This prevents an API request for
     * every single keyboard press.
     */
    const timer = setTimeout(async () => {
      try {
        setSearchLoading(true);
        setError("");

        /*
         * ======================================
         * EXACT PLAYER ID SEARCH
         * ======================================
         *
         * Example:
         * aardsda01
         *
         * If it looks like a complete ID,
         * try the ID API first.
         */

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
            /*
             * If exact ID doesn't exist,
             * continue with name search.
             */

            if (err.response?.status !== 404) {
              throw err;
            }
          }
        }

        /*
         * ======================================
         * NAME SEARCH
         * ======================================
         *
         * Example:
         * mi
         * mike
         * michael
         */

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

    /*
     * If the user types again before 300ms,
     * cancel the previous search.
     */
    return () => clearTimeout(timer);
  }, [searchText]);

  /*
   * ==========================================
   * SEARCH INPUT
   * ==========================================
   */

  const handleSearchChange = (event) => {
    const value = event.target.value;

    setSearchText(value);

    /*
     * When user starts searching,
     * reset pagination to page 1.
     */
    setPage(1);

    /*
     * Clear previous errors.
     */
    setError("");
  };

  /*
   * ==========================================
   * CLEAR SEARCH
   * ==========================================
   */

  const handleClearSearch = () => {
    setSearchText("");
    setSearchResults([]);
    setError("");
    setPage(1);
  };

  /*
   * ==========================================
   * GET PERSON DETAILS
   * ==========================================
   */

  const handlePersonClick = async (playerID) => {
    try {
      setLoading(true);
      setError("");

      const person = await getPersonById(playerID);

      setSelectedPerson(person);
    } catch (err) {
      console.error(err);
      setError("Failed to load person");
    } finally {
      setLoading(false);
    }
  };

  /*
   * ==========================================
   * BACK TO PEOPLE TABLE
   * ==========================================
   */

  const handleBackToPeople = () => {
    setSelectedPerson(null);
    setError("");
  };

  /*
   * ==========================================
   * DETAILS PAGE
   * ==========================================
   */

  if (selectedPerson) {
    return (
      <PersonDetails
        person={selectedPerson}
        onClose={handleBackToPeople}
      />
    );
  }

  /*
   * ==========================================
   * WHICH DATA SHOULD TABLE SHOW?
   * ==========================================
   *
   * Searching:
   *     searchResults
   *
   * Normal:
   *     people
   */

  const tableData =
    searchText.trim().length >= 2
      ? searchResults
      : people;

  const isSearching =
    searchText.trim().length >= 2;

  return (
    <div className="app">
      <div className="container">

        {/* =================================
            HEADER
        ================================= */}

        <div className="header">
          <div>
            <h1>People API</h1>

            <p>
              Browse and explore player information
            </p>
          </div>
        </div>

        {/* =================================
            SEARCH
        ================================= */}

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

        </div>

        {/* =================================
            ERROR
        ================================= */}

        {error && (
          <div className="error">
            {error}
          </div>
        )}

        {/* =================================
            TABLE / LOADING
        ================================= */}

        {loading || searchLoading ? (
          <div className="loading">
            {searchLoading
              ? "Searching..."
              : "Loading..."}
          </div>
        ) : (
          <>
            {/*
             * =================================
             * SEARCH RESULTS
             *
             * The same PeopleTable is used.
             * Therefore search results look
             * exactly like normal table data.
             * =================================
             */}

            <PeopleTable
              people={tableData}
              onPersonClick={handlePersonClick}
            />

            {/*
             * =================================
             * PAGINATION
             *
             * Don't show pagination while
             * searching.
             * =================================
             */}

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

export default App;