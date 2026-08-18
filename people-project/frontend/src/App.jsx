import { useEffect, useState } from "react";
import "./App.css";

import { getPeople, getPersonById } from "./api/peopleApi";

import PeopleTable from "./components/PeopleTable";
import Pagination from "./components/Pagination";
import PersonDetails from "./components/PersonDetails";

function App() {
  const [people, setPeople] = useState([]);

  const [page, setPage] = useState(1);
  const [limit] = useState(10);
  const [totalPages, setTotalPages] = useState(1);

  const [selectedPerson, setSelectedPerson] = useState(null);

  const [searchId, setSearchId] = useState("");

  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");

  useEffect(() => {
    loadPeople();
  }, [page]);

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

  const handleSearch = async (event) => {
    event.preventDefault();

    const playerID = searchId.trim();

    if (!playerID) {
      setError("Please enter a Player ID");
      return;
    }

    try {
      setLoading(true);
      setError("");
      setSelectedPerson(null);

      const person = await getPersonById(playerID);

      setSelectedPerson(person);
    } catch (err) {
      console.error(err);

      if (err.response?.status === 404) {
        setError(`Player "${playerID}" not found`);
      } else {
        setError("Failed to search player");
      }
    } finally {
      setLoading(false);
    }
  };

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

  const handleBackToPeople = () => {
    setSelectedPerson(null);
    setError("");
  };

  if (selectedPerson) {
    return (
      <PersonDetails
        person={selectedPerson}
        onClose={handleBackToPeople}
      />
    );
  }

  return (
    <div className="app">
      <div className="container">

        {/* Header */}

        <div className="header">
          <div>
            <h1>People API</h1>

            <p>
              Browse and explore player information
            </p>
          </div>
        </div>

        {/* Search */}

        <form
          className="search-bar"
          onSubmit={handleSearch}
        >
          <input
            type="text"
            placeholder="Search by Player ID..."
            value={searchId}
            onChange={(event) =>
              setSearchId(event.target.value)
            }
          />

          <button type="submit">
            Search
          </button>
        </form>

        {/* Error */}

        {error && (
          <div className="error">
            {error}
          </div>
        )}

        {/* People */}

        {loading ? (
          <div className="loading">
            Loading...
          </div>
        ) : (
          <>
            <PeopleTable
              people={people}
              onPersonClick={handlePersonClick}
            />

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
          </>
        )}

      </div>
    </div>
  );
}

export default App;