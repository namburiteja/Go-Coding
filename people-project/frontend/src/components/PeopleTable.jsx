import { useEffect, useRef } from "react";

function PeopleTable({
  people,
  onPersonClick,
  infiniteScroll = false,
  hasNext = false,
  loadingMore = false,
  onLoadMore,
}) {
  const wrapperRef = useRef(null);
  const sentinelRef = useRef(null);

  useEffect(() => {
    if (!infiniteScroll) {
      return;
    }

    const sentinel = sentinelRef.current;
    const wrapper = wrapperRef.current;

    if (!sentinel || !wrapper) {
      return;
    }

    const observer = new IntersectionObserver(
      (entries) => {
        const entry = entries[0];

        if (
          entry.isIntersecting &&
          hasNext &&
          !loadingMore
        ) {
          onLoadMore();
        }
      },
      {
        root: wrapper,
        rootMargin: "150px",
      }
    );

    observer.observe(sentinel);

    return () => {
      observer.disconnect();
    };
  }, [
    infiniteScroll,
    hasNext,
    loadingMore,
    onLoadMore,
  ]);

  return (
    <div className="table-card">
      <div
        className="table-wrapper"
        ref={wrapperRef}
      >
        <table className="people-table">
          <thead>
            <tr>
              <th>Player ID</th>
              <th>First Name</th>
              <th>Last Name</th>
              <th>Birth Year</th>
              <th>Birth Country</th>
              <th>Birth City</th>
              <th>Height</th>
            </tr>
          </thead>

          <tbody>
            {people.map((person) => (
              <tr
                key={person.playerID}
                className="people-row"
                onClick={() =>
                  onPersonClick(person.playerID)
                }
              >
                <td className="player-id">
                  {person.playerID}
                </td>

                <td className="name">
                  {person.nameFirst ?? "-"}
                </td>

                <td className="name">
                  {person.nameLast ?? "-"}
                </td>

                <td>
                  {person.birthYear ?? "-"}
                </td>

                <td>
                  {person.birthCountry ?? "-"}
                </td>

                <td>
                  {person.birthCity ?? "-"}
                </td>

                <td>
                  {person.height ?? "-"}
                </td>
              </tr>
            ))}
          </tbody>
        </table>

        {infiniteScroll && (
          <div
            ref={sentinelRef}
            className="infinite-scroll-sentinel"
          >
            {loadingMore && "Loading more..."}
          </div>
        )}
      </div>
    </div>
  );
}

export default PeopleTable;