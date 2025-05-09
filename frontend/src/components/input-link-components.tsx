import { useState, type ChangeEvent } from "react";

type Props = {
  input: { url: string };
  setInput: React.Dispatch<React.SetStateAction<{ url: string }>>;
};

const InputLink = ({ input, setInput }: Props) => {
  const [url, setUrl] = useState<string>("");

  const handleClick = () => {
    setInput({ url });
  };

  return (
    <div className="bg-purple-600 p-4 rounded-xl w-xl border-amber-500">
      <div className="flex flex-col my-4">
        <label htmlFor="url" className="text-4xl my-2 font-bold font-sans ">
          Title:
        </label>
        <input
          id="title"
          type="text"
          name="url"
          onChange={(e: ChangeEvent<HTMLInputElement>) =>
            setUrl(e.target.value)
          }
          className="bg-gray-200 p-2 rounded-md focus:bg-gray-400 text-black"
        />
      </div>
      <button
        onClick={handleClick}
        className="p-2 w-full text-xl bg-gray-400 hover:bg-gray-500 rounded-2xl cursor-pointer font-black font-sans "
      >
        Short link
      </button>
    </div>
  );
};

export default InputLink;
