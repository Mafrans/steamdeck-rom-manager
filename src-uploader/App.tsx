import { HashRouter, Route, Routes } from "react-router-dom";
import { HomeView } from "./views/HomeView/HomeView";

export function App() {
  return (
    <>
      <HashRouter>
        <Routes>
          <Route path="/" element={<HomeView />} />
        </Routes>
      </HashRouter>
    </>
  );
}
