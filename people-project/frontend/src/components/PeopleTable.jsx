function PeopleTable({ people, onPersonClick }) {
  return (
    <div className="table-card">
      <div className="table-wrapper">
        <table className="people-table">
          <thead>
            <tr>
              <th>Player ID</th>
              <th>First Name</th>
              <th>Last Name</th>
              <th>Birth Year</th>
              <th>Birth Country</th>
              <th>Birth City</th>
            </tr>
          </thead>

          <tbody>
            {people.map((person) => (
              <tr
                key={person.playerID}
                className="people-row"
                onClick={() => onPersonClick(person.playerID)}
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
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
}

export default PeopleTable;