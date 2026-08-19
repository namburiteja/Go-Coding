import axios from "axios";

const API = axios.create({
  baseURL: "http://localhost:8090",
});

export const getPeople = async (page = 1, limit = 10) => {
  const response = await API.get("/people", {
    params: {
      page,
      limit,
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