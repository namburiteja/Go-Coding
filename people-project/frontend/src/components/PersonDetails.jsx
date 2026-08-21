import { useState } from "react";
import { updatePerson } from "../api/peopleApi";

function PersonDetails({ person, onClose }) {
  const [isEditing, setIsEditing] = useState(false);
  const [saving, setSaving] = useState(false);
  const [error, setError] = useState("");

  const [formData, setFormData] = useState({
    nameFirst: person.nameFirst ?? "",
    nameLast: person.nameLast ?? "",
    nameGiven: person.nameGiven ?? "",

    birthYear: person.birthYear ?? "",
    birthMonth: person.birthMonth ?? "",
    birthDay: person.birthDay ?? "",
    birthCountry: person.birthCountry ?? "",
    birthState: person.birthState ?? "",
    birthCity: person.birthCity ?? "",

    deathYear: person.deathYear ?? "",
    deathMonth: person.deathMonth ?? "",
    deathDay: person.deathDay ?? "",
    deathCountry: person.deathCountry ?? "",
    deathState: person.deathState ?? "",
    deathCity: person.deathCity ?? "",

    weight: person.weight ?? "",
    height: person.height ?? "",

    bats: person.bats ?? "",
    throws: person.throws ?? "",

    debut: person.debut
      ? person.debut.split("T")[0]
      : "",

    finalGame: person.finalGame
      ? person.finalGame.split("T")[0]
      : "",

    retroID: person.retroID ?? "",
    bbrefID: person.bbrefID ?? "",
  });

  if (!person) {
    return null;
  }

  const handleChange = (event) => {
    const { name, value } = event.target;

    setFormData((previous) => ({
      ...previous,
      [name]: value,
    }));
  };

  const handleEdit = () => {
    setError("");
    setIsEditing(true);
  };

  const handleCancel = () => {
    setFormData({
      nameFirst: person.nameFirst ?? "",
      nameLast: person.nameLast ?? "",
      nameGiven: person.nameGiven ?? "",

      birthYear: person.birthYear ?? "",
      birthMonth: person.birthMonth ?? "",
      birthDay: person.birthDay ?? "",
      birthCountry: person.birthCountry ?? "",
      birthState: person.birthState ?? "",
      birthCity: person.birthCity ?? "",

      deathYear: person.deathYear ?? "",
      deathMonth: person.deathMonth ?? "",
      deathDay: person.deathDay ?? "",
      deathCountry: person.deathCountry ?? "",
      deathState: person.deathState ?? "",
      deathCity: person.deathCity ?? "",

      weight: person.weight ?? "",
      height: person.height ?? "",

      bats: person.bats ?? "",
      throws: person.throws ?? "",

      debut: person.debut
        ? person.debut.split("T")[0]
        : "",

      finalGame: person.finalGame
        ? person.finalGame.split("T")[0]
        : "",

      retroID: person.retroID ?? "",
      bbrefID: person.bbrefID ?? "",
    });

    setError("");
    setIsEditing(false);
  };

  const handleSave = async () => {
    try {
      setSaving(true);
      setError("");

      const payload = {
        nameFirst: formData.nameFirst,
        nameLast: formData.nameLast,
        nameGiven: formData.nameGiven,

        birthYear:
          formData.birthYear === ""
            ? null
            : Number(formData.birthYear),

        birthMonth:
          formData.birthMonth === ""
            ? null
            : Number(formData.birthMonth),

        birthDay:
          formData.birthDay === ""
            ? null
            : Number(formData.birthDay),

        birthCountry: formData.birthCountry,
        birthState: formData.birthState,
        birthCity: formData.birthCity,

        deathYear:
          formData.deathYear === ""
            ? null
            : Number(formData.deathYear),

        deathMonth:
          formData.deathMonth === ""
            ? null
            : Number(formData.deathMonth),

        deathDay:
          formData.deathDay === ""
            ? null
            : Number(formData.deathDay),

        deathCountry: formData.deathCountry,
        deathState: formData.deathState,
        deathCity: formData.deathCity,

        weight:
          formData.weight === ""
            ? null
            : Number(formData.weight),

        height:
          formData.height === ""
            ? null
            : Number(formData.height),

        bats: formData.bats,
        throws: formData.throws,

        debut: formData.debut,
        finalGame: formData.finalGame,

        retroID: formData.retroID,
        bbrefID: formData.bbrefID,
      };

      await updatePerson(person.playerID, payload);

      setIsEditing(false);

      alert("Person updated successfully!");

      window.location.reload();
    } catch (err) {
      console.error("Update failed:", err);

      setError(
        err.response?.data ||
          "Failed to update person"
      );
    } finally {
      setSaving(false);
    }
  };

  const fields = [
    ["First Name", "nameFirst", "text"],
    ["Last Name", "nameLast", "text"],
    ["Given Name", "nameGiven", "text"],

    ["Birth Year", "birthYear", "number"],
    ["Birth Month", "birthMonth", "number"],
    ["Birth Day", "birthDay", "number"],
    ["Birth Country", "birthCountry", "text"],
    ["Birth State", "birthState", "text"],
    ["Birth City", "birthCity", "text"],

    ["Death Year", "deathYear", "number"],
    ["Death Month", "deathMonth", "number"],
    ["Death Day", "deathDay", "number"],
    ["Death Country", "deathCountry", "text"],
    ["Death State", "deathState", "text"],
    ["Death City", "deathCity", "text"],

    ["Weight", "weight", "number"],
    ["Height", "height", "number"],

    ["Bats", "bats", "text"],
    ["Throws", "throws", "text"],

    ["Debut", "debut", "date"],
    ["Final Game", "finalGame", "date"],

    ["Retro ID", "retroID", "text"],
    ["BBRef ID", "bbrefID", "text"],
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

            {!isEditing && (
              <button
                className="edit-button"
                onClick={handleEdit}
              >
                Edit Person
              </button>
            )}

          </div>

          {error && (
            <div className="error">
              {error}
            </div>
          )}

          <div className="details-grid">

            {/* PLAYER ID - NEVER EDITABLE */}
            <div className="detail-item">
              <span className="detail-label">
                Player ID
              </span>

              <span className="detail-value">
                {person.playerID}
              </span>
            </div>

            {fields.map(([label, name, type]) => (
              <div
                className="detail-item"
                key={name}
              >
                <span className="detail-label">
                  {label}
                </span>

                {isEditing ? (
                  <input
                    className="detail-input"
                    type={type}
                    name={name}
                    value={formData[name]}
                    onChange={handleChange}
                  />
                ) : (
                  <span className="detail-value">
                    {formData[name] || "-"}
                  </span>
                )}
              </div>
            ))}

          </div>

          {isEditing && (
            <div className="edit-actions">

              <button
                className="cancel-button"
                onClick={handleCancel}
                disabled={saving}
              >
                Cancel
              </button>

              <button
                className="save-button"
                onClick={handleSave}
                disabled={saving}
              >
                {saving
                  ? "Saving..."
                  : "Save Changes"}
              </button>

            </div>
          )}

        </div>
      </div>
    </div>
  );
}

export default PersonDetails;