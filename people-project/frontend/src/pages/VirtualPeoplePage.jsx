import {
  useCallback,
  useEffect,
  useState,
} from "react";

import "../App.css";

import {
  getPeopleCursor,
  getPersonById,
  searchPeopleByName,
} from "../api/peopleApi";

import VirtualPeopleTable from "../components/VirtualPeopleTable";

function VirtualPeoplePage({ onPersonClick }) {
  /*
   * ==========================================
   * CURSOR PAGINATION DATA
   * ==========================================
   */

  const [people, setPeople] = useState([]);

  const [nextCursor, setNextCursor] = useState("");

  const [hasNext, setHasNext] = useState(true);

  const [loading, setLoading] = useState(false);

  const [loadingMore, setLoadingMore] =
    useState(false);

  /*
   * ==========================================
   * SEARCH
   * ==========================================
   */

  const [searchText, setSearchText] =
    useState("");

  const [searchResults, setSearchResults] =
    useState([]);

  const [searchLoading, setSearchLoading] =
    useState(false);

  /*
   * ==========================================
   * PAGE SIZE
   * ==========================================
   */

  const limit = 100;

  /*
   * ==========================================
   * INITIAL CURSOR LOAD
   * ==========================================
   */

  const loadInitialPeople =
    useCallback(async () => {
      try {
        setLoading(true);

        const data = await getPeopleCursor(
          "",
          limit
        );

        setPeople(data.data);

        setNextCursor(
          data.nextCursor || ""
        );

        setHasNext(data.hasNext);
      } catch (err) {
        console.error(err);
      } finally {
        setLoading(false);
      }
    }, []);

  /*
   * ==========================================
   * LOAD INITIAL DATA
   * ==========================================
   */

  useEffect(() => {
    loadInitialPeople();
  }, [loadInitialPeople]);

  /*
   * ==========================================
   * LOAD MORE CURSOR DATA
   * ==========================================
   */

  const loadMorePeople = useCallback(
    async () => {
      /*
       * Don't load more while searching.
       */

      if (
        searchText.trim() !== "" ||
        !hasNext ||
        loadingMore ||
        !nextCursor
      ) {
        return;
      }

      try {
        setLoadingMore(true);

        const data =
          await getPeopleCursor(
            nextCursor,
            limit
          );

        /*
         * IMPORTANT:
         *
         * Cursor pagination keeps
         * previous records.
         *
         * New records are appended.
         */

        setPeople(
          (currentPeople) => [
            ...currentPeople,
            ...data.data,
          ]
        );

        setNextCursor(
          data.nextCursor || ""
        );

        setHasNext(data.hasNext);
      } catch (err) {
        console.error(err);
      } finally {
        setLoadingMore(false);
      }
    },
    [
      searchText,
      hasNext,
      loadingMore,
      nextCursor,
    ]
  );

  /*
   * ==========================================
   * VIRTUALIZED LIST
   * ==========================================
   *
   * When the virtualized list gets close
   * to the end, request the next cursor page.
   */

  const handleItemsRendered = ({
    stopIndex,
  }) => {
    /*
     * Don't load more while searching.
     */

    if (searchText.trim() !== "") {
      return;
    }

    const threshold = 10;

    if (
      stopIndex >=
        people.length - threshold &&
      hasNext &&
      !loadingMore
    ) {
      loadMorePeople();
    }
  };

  /*
   * ==========================================
   * SEARCH
   * ==========================================
   */

  useEffect(() => {
    const value = searchText.trim();

    /*
     * Empty search
     *
     * Return to cursor data.
     */

    if (value === "") {
      setSearchResults([]);
      setSearchLoading(false);
      return;
    }

    /*
     * Don't search until 2 characters.
     */

    if (value.length < 2) {
      setSearchResults([]);
      setSearchLoading(false);
      return;
    }

    /*
     * Debounce search.
     *
     * Wait 300ms after user stops typing.
     */

    const timer = setTimeout(
      async () => {
        try {
          setSearchLoading(true);

          /*
           * ==================================
           * EXACT PLAYER ID SEARCH
           * ==================================
           */

          const looksLikePlayerID =
            /^[a-zA-Z0-9]+$/.test(value) &&
            /\d/.test(value) &&
            value.length >= 6;

          if (looksLikePlayerID) {
            try {
              const person =
                await getPersonById(value);

              setSearchResults([
                person,
              ]);

              return;
            } catch (err) {
              /*
               * If ID doesn't exist,
               * continue with name search.
               */

              if (
                err.response?.status !==
                404
              ) {
                throw err;
              }
            }
          }

          /*
           * ==================================
           * NAME SEARCH
           * ==================================
           */

          const results =
            await searchPeopleByName(
              value
            );

          setSearchResults(results);
        } catch (err) {
          console.error(err);

          setSearchResults([]);
        } finally {
          setSearchLoading(false);
        }
      },
      300
    );

    return () =>
      clearTimeout(timer);
  }, [searchText]);

  /*
   * ==========================================
   * SEARCH INPUT
   * ==========================================
   */

  const handleSearchChange = (
    event
  ) => {
    setSearchText(
      event.target.value
    );
  };

  /*
   * ==========================================
   * CLEAR SEARCH
   * ==========================================
   */

  const handleClearSearch = () => {
    setSearchText("");

    setSearchResults([]);
  };

  /*
   * ==========================================
   * WHICH DATA SHOULD BE DISPLAYED?
   * ==========================================
   *
   * Searching:
   *
   *     searchResults
   *
   * Normal:
   *
   *     people
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

        {/* =================================
            HEADER
        ================================= */}

        <div className="header">
          <div>
            <h1>People API</h1>

            <p>
              Cursor Pagination •
              Virtualized List
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
              onChange={
                handleSearchChange
              }
            />

            {searchText && (
              <button
                type="button"
                className="search-clear"
                onClick={
                  handleClearSearch
                }
                aria-label="Clear search"
              >
                ×
              </button>
            )}

          </div>

        </div>

        {/* =================================
            LOADING
        ================================= */}

        {loading ? (
          <div className="loading">
            Loading...
          </div>
        ) : searchLoading ? (
          <div className="loading">
            Searching...
          </div>
        ) : (
          <>

            {/* =================================
                VIRTUALIZED TABLE
            ================================= */}

            <VirtualPeopleTable
              people={tableData}
              onPersonClick={
                onPersonClick
              }
              onItemsRendered={
                isSearching
                  ? undefined
                  : handleItemsRendered
              }
            />

            {/* =================================
                LOADING MORE
            ================================= */}

            {!isSearching &&
              loadingMore && (
                <div className="loading">
                  Loading more...
                </div>
              )}

            {/* =================================
                END OF DATA
            ================================= */}

            {!isSearching &&
              !hasNext && (
                <div className="loading">
                  No more people
                </div>
              )}

          </>
        )}

      </div>
    </div>
  );
}

export default VirtualPeoplePage;