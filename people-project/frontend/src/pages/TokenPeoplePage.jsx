import { useCallback, useEffect, useState } from "react";
import "../App.css";

import {
  getPeopleToken,
} from "../api/peopleApi";

import PeopleTable from "../components/PeopleTable";

function TokenPeoplePage({ onPersonClick }) {
  const [people, setPeople] = useState([]);

  const [nextToken, setNextToken] = useState("");
  const [hasNext, setHasNext] = useState(true);

  const [loading, setLoading] = useState(false);
  const [loadingMore, setLoadingMore] = useState(false);

  const [error, setError] = useState("");

  const limit = 20;

  /*
   * ==========================================
   * INITIAL LOAD
   * ==========================================
   */

  const loadInitialPeople = useCallback(async () => {
    try {
      setLoading(true);
      setError("");

      const data = await getPeopleToken("", limit);

      setPeople(data.data);

      setNextToken(data.nextToken || "");

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

  const loadMorePeople = async () => {
    if (
      !hasNext ||
      loadingMore ||
      !nextToken
    ) {
      return;
    }

    try {
      setLoadingMore(true);
      setError("");

      const data = await getPeopleToken(
        nextToken,
        limit
      );

      setPeople((currentPeople) => [
        ...currentPeople,
        ...data.data,
      ]);

      setNextToken(
        data.nextToken || ""
      );

      setHasNext(data.hasNext);
    } catch (err) {
      console.error(err);
      setError("Failed to load more people");
    } finally {
      setLoadingMore(false);
    }
  };

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
              Token Pagination • Load More
            </p>
          </div>
        </div>

        {error && (
          <div className="error">
            {error}
          </div>
        )}

        {loading ? (
          <div className="loading">
            Loading...
          </div>
        ) : (
          <>
            <PeopleTable
              people={people}
              onPersonClick={onPersonClick}
            />

            <div className="load-more-container">

              {hasNext ? (
                <button
                  className="load-more-button"
                  onClick={loadMorePeople}
                  disabled={loadingMore}
                >
                  {loadingMore
                    ? "Loading..."
                    : "Load More"}
                </button>
              ) : (
                <div className="loading">
                  No more people
                </div>
              )}

            </div>
          </>
        )}

      </div>
    </div>
  );
}

export default TokenPeoplePage;