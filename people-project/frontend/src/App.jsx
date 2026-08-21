import { useState } from "react";

import {
  getPersonById,
} from "./api/peopleApi";

import OffsetPeoplePage from "./pages/OffsetPeoplePage";
import CursorPeoplePage from "./pages/CursorPeoplePage";
import PersonDetails from "./components/PersonDetails";
import VirtualPeoplePage from "./pages/VirtualPeoplePage";
import TokenPeoplePage from "./pages/TokenPeoplePage";

function App() {
  const [selectedPerson, setSelectedPerson] = useState(null);

  const path = window.location.pathname;


  //  GET PERSON DETAILS

  const handlePersonClick = async (playerID) => {
    try {
      const person = await getPersonById(playerID);

      setSelectedPerson(person);
    } catch (err) {
      console.error("Failed to load person:", err);
    }
  };

  //  BACK TO PEOPLE

  const handleBackToPeople = () => {
    setSelectedPerson(null);
  };

  
   // PERSON DETAILS
   

  if (selectedPerson) {
    return (
      <PersonDetails
        person={selectedPerson}
        onClose={handleBackToPeople}
      />
    );
  }

  //  VIRTUALIZATION PAGE

  if (path === "/virtual") {
  return (
    <VirtualPeoplePage
      onPersonClick={handlePersonClick}
    />
  );
}


 // TOKEN PAGE
  if (path === "/token") {
    return (
      <TokenPeoplePage
        onPersonClick={handlePersonClick}
      />
    );
  }


   // CURSOR PAGE


  if (path === "/cursor") {
    return (
      <CursorPeoplePage
        onPersonClick={handlePersonClick}
      />
    );
  }

  //  OFFSET PAGE

  return (
    <OffsetPeoplePage
      onPersonClick={handlePersonClick}
    />
  );
}

export default App;