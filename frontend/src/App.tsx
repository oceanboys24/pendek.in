import { useState } from "react";
import "./App.css";
import InputLink from "./components/input-link-components";
import ResultComponents from "./components/result-components";

function App() {
  const [input, setInput] = useState<{ url: string }>({
    url: "",
  });
  return (
    <div className="bg-gray-300 w-full h-screen">
      <div className="flex flex-col justify-center items-center h-full gap-4 text-white">
        <InputLink input={input} setInput={setInput} />
        <ResultComponents url={input.url} />
      </div>
    </div>
  );
}

export default App;
