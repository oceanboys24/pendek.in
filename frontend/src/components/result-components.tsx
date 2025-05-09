const ResultComponents = ({ url }: { url: string }) => {
  return (
    <div className="bg-purple-600 p-4 rounded-xl w-xl text-black  font-serif">
      <h1 className="font-bold text-3xl">Result:</h1>
      <div className="flex gap-4">
        <h1>Link:</h1>
        <h1>{url}</h1>
      </div>
    </div>
  );
};

export default ResultComponents;
