import { List } from "react-window";

function Row({ index, style, people, onPersonClick }) {
  const person = people[index];

  return (
    <div
      style={style}
      className="virtual-row people-row"
      onClick={() =>
        onPersonClick(person.playerID)
      }
    >
      <div className="virtual-cell player-id">
        {person.playerID}
      </div>

      <div className="virtual-cell name">
        {person.nameFirst ?? "-"}
      </div>

      <div className="virtual-cell name">
        {person.nameLast ?? "-"}
      </div>

      <div className="virtual-cell">
        {person.birthYear ?? "-"}
      </div>

      <div className="virtual-cell">
        {person.birthCountry ?? "-"}
      </div>

      <div className="virtual-cell">
        {person.birthCity ?? "-"}
      </div>
    </div>
  );
}

function VirtualPeopleTable({
  people,
  onPersonClick,
  onItemsRendered,
}) {
  return (
    <div className="table-card virtual-table-card">

      <div className="virtual-header">
        <div>Player ID</div>
        <div>First Name</div>
        <div>Last Name</div>
        <div>Birth Year</div>
        <div>Birth Country</div>
        <div>Birth City</div>
      </div>

      <List
        rowComponent={Row}
        rowCount={people.length}
        rowHeight={52}
        rowProps={{
          people,
          onPersonClick,
        }}
        defaultHeight={500}
        onRowsRendered={onItemsRendered}
        overscanCount={5}
        style={{
          height: "500px",
          width: "100%",
        }}
      />

    </div>
  );
}

export default VirtualPeopleTable;