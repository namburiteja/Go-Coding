function PersonDetails({ person, onClose }) {
  if (!person) {
    return null;
  }

  const formatDate = (date) => {
    if (!date) {
      return "-";
    }

    return date.split("T")[0];
  };

  const fields = [
    ["Player ID", person.playerID],
    ["First Name", person.nameFirst],
    ["Last Name", person.nameLast],
    ["Given Name", person.nameGiven],

    ["Birth Year", person.birthYear],
    ["Birth Month", person.birthMonth],
    ["Birth Day", person.birthDay],
    ["Birth Country", person.birthCountry],
    ["Birth State", person.birthState],
    ["Birth City", person.birthCity],

    ["Death Year", person.deathYear],
    ["Death Month", person.deathMonth],
    ["Death Day", person.deathDay],
    ["Death Country", person.deathCountry],
    ["Death State", person.deathState],
    ["Death City", person.deathCity],

    ["Weight", person.weight],
    ["Height", person.height],
    ["Bats", person.bats],
    ["Throws", person.throws],

    ["Debut", formatDate(person.debut)],
    ["Final Game", formatDate(person.finalGame)],

    ["Retro ID", person.retroID],
    ["BBRef ID", person.bbrefID],
  ];

  return (
    <div className="details-page">
      <div className="details-container">

        <button
          className="back-button"
          onClick={onClose}
        >
          ← Back to People
        </button>

        <div className="details-card">

          <div className="details-header">
            <div>
              <span className="details-label">
                PLAYER DETAILS
              </span>

              <h1>
                {person.nameFirst ?? ""}{" "}
                {person.nameLast ?? ""}
              </h1>

              <p>
                Player ID: {person.playerID}
              </p>
            </div>
          </div>

          <div className="details-grid">
            {fields.map(([label, value]) => (
              <div
                className="detail-item"
                key={label}
              >
                <span className="detail-label">
                  {label}
                </span>

                <span className="detail-value">
                  {value ?? "-"}
                </span>
              </div>
            ))}
          </div>

        </div>

      </div>
    </div>
  );
}

export default PersonDetails;