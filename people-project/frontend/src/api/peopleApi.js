import axios from "axios";

const API = axios.create({
  baseURL: "http://localhost:8090",
});

export const getPeople = async (
  page = 1,
  limit = 10,
  sortBy = "",
  sortOrder = "",
  filters = {}
) => {
  const response = await API.get("/people", {
    params: {
      page,
      limit,

      ...(sortBy && { sortBy }),
      ...(sortOrder && { sortOrder }),

      ...(filters.birthCountry && {
        birthCountry: filters.birthCountry,
      }),

      ...(filters.birthYearFrom && {
        birthYearFrom: filters.birthYearFrom,
      }),

      ...(filters.birthYearTo && {
        birthYearTo: filters.birthYearTo,
      }),

      ...(filters.heightFrom && {
        heightFrom: filters.heightFrom,
      }),

      ...(filters.heightTo && {
        heightTo: filters.heightTo,
      }),

      ...(filters.weightFrom && {
        weightFrom: filters.weightFrom,
      }),

      ...(filters.weightTo && {
        weightTo: filters.weightTo,
      }),

      ...(filters.bats && {
        bats: filters.bats,
      }),

      ...(filters.throws && {
        throws: filters.throws,
      }),
    },
  });

  return response.data;
};

export const getPersonById = async (playerID) => {
  const response = await API.get(`/people/${playerID}`);
  return response.data;
};

export const searchPeopleByName = async (name) => {
  const response = await API.get("/people/search", {
    params: {
      name,
    },
  });

  return response.data;
};

export const getPeopleCursor = async (cursor = "", limit = 10) => {
  const response = await API.get("/people/cursor", {
    params: {
      limit,
      ...(cursor && { cursor }),
    },
  });

  return response.data;
};

export const getPeopleToken = async (token = "", limit = 20) => {
  const response = await API.get("/people/token", {
    params: {
      limit,
      ...(token && { token }),
    },
  });

  return response.data;
};

// UPDATE PERSON
export const updatePerson = async (playerID, person) => {
  const response = await API.put("/people/update", person, {
    params: {
      playerID,
    },
  });

  return response.data;
};