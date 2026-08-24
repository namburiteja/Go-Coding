import { useState } from "react";

function PeopleFilters({ onApply, onClear }) {
  const [filters, setFilters] = useState({
    birthCountry: "",
    birthYearFrom: "",
    birthYearTo: "",
    heightFrom: "",
    heightTo: "",
    weightFrom: "",
    weightTo: "",
    bats: "",
    throws: "",
  });

  const handleChange = (event) => {
    const { name, value } = event.target;

    setFilters((previous) => ({
      ...previous,
      [name]: value,
    }));
  };

  const handleApply = () => {
    onApply(filters);
  };

  const handleClear = () => {
    const emptyFilters = {
      birthCountry: "",
      birthYearFrom: "",
      birthYearTo: "",
      heightFrom: "",
      heightTo: "",
      weightFrom: "",
      weightTo: "",
      bats: "",
      throws: "",
    };

    setFilters(emptyFilters);

    onClear();
  };

  return (
    <div className="filters-panel">

      <h3>Filters</h3>

      <div className="filters-grid">

        {/* Birth Country */}
        <div className="filter-group">
          <label>Birth Country</label>

          <input
            type="text"
            name="birthCountry"
            placeholder="e.g. USA"
            value={filters.birthCountry}
            onChange={handleChange}
          />
        </div>

        {/* Birth Year From */}
        <div className="filter-group">
          <label>Birth Year From</label>

          <input
            type="number"
            name="birthYearFrom"
            placeholder="e.g. 1900"
            value={filters.birthYearFrom}
            onChange={handleChange}
          />
        </div>

        {/* Birth Year To */}
        <div className="filter-group">
          <label>Birth Year To</label>

          <input
            type="number"
            name="birthYearTo"
            placeholder="e.g. 2000"
            value={filters.birthYearTo}
            onChange={handleChange}
          />
        </div>

        {/* Height From */}
        <div className="filter-group">
          <label>Height From</label>

          <input
            type="number"
            name="heightFrom"
            placeholder="e.g. 160"
            value={filters.heightFrom}
            onChange={handleChange}
          />
        </div>

        {/* Height To */}
        <div className="filter-group">
          <label>Height To</label>

          <input
            type="number"
            name="heightTo"
            placeholder="e.g. 200"
            value={filters.heightTo}
            onChange={handleChange}
          />
        </div>

        {/* Weight From */}
        <div className="filter-group">
          <label>Weight From</label>

          <input
            type="number"
            name="weightFrom"
            placeholder="e.g. 50"
            value={filters.weightFrom}
            onChange={handleChange}
          />
        </div>

        {/* Weight To */}
        <div className="filter-group">
          <label>Weight To</label>

          <input
            type="number"
            name="weightTo"
            placeholder="e.g. 100"
            value={filters.weightTo}
            onChange={handleChange}
          />
        </div>

        {/* Bats */}
        <div className="filter-group">
          <label>Bats</label>

          <select
            name="bats"
            value={filters.bats}
            onChange={handleChange}
          >
            <option value="">Any</option>
            <option value="R">Right</option>
            <option value="L">Left</option>
            <option value="B">Both</option>
          </select>
        </div>

        {/* Throws */}
        <div className="filter-group">
          <label>Throws</label>

          <select
            name="throws"
            value={filters.throws}
            onChange={handleChange}
          >
            <option value="">Any</option>
            <option value="R">Right</option>
            <option value="L">Left</option>
          </select>
        </div>

      </div>

      <div className="filter-actions">

        <button
          type="button"
          onClick={handleApply}
        >
          Apply Filters
        </button>

        <button
          type="button"
          onClick={handleClear}
        >
          Clear Filters
        </button>

      </div>

    </div>
  );
}

export default PeopleFilters;