import { useState } from "react";
import { useInputUrl } from "../hook/useInputUrl";
import { useShortLink } from "../hook/useShortLink";
import ResultComponents from "./result-components";

const InputLink = () => {
  const { register, handleSubmit, reset } = useInputUrl();
  const { mutateAsync, isPending } = useShortLink();
  const [url, setUrl] = useState<string>("");

  const onsubmit = async (url: { url: string }) => {
    const res = mutateAsync(url);
    reset({
      url: "",
    });
    setUrl((await res).short_link);
  };

  return (
    <div className="flex flex-col justify-center items-center h-full gap-4 text-white">
      <div className="bg-purple-600 p-4 rounded-xl w-xl border-amber-500">
        <div className="flex flex-col my-4">
          <label htmlFor="url" className="text-4xl my-2 font-bold font-sans ">
            Title:
          </label>
          <input
            id="title"
            type="text"
            placeholder="http://localhost:5173/"
            {...register("url")}
            className="bg-gray-200 p-2 rounded-md focus:bg-gray-400 text-black"
          />
        </div>
        <button
          onClick={handleSubmit(onsubmit)}
          className="p-2 w-full text-xl bg-gray-400 hover:bg-gray-500 rounded-2xl cursor-pointer font-black font-sans "
        >
          {isPending ? "Loading..." : "Short link"}
        </button>
      </div>
      <ResultComponents url={url} />
    </div>
  );
};

export default InputLink;
