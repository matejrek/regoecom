import { useState, useMemo } from "react";
import ProductBox from "./productBox";
import useFetch from "../hooks/useFetch";

function Search() {
  const [term, setTerm] = useState("");
  const [data, error] = useFetch('products');
  const [result, setResult] = useState<any | null>(null);

  // Use useMemo to memoize the result state
  const memoizedResult = useMemo(() => {
    if (term === '') {
      return null;
    }
    return data.filter((data: any) => JSON.stringify(data.title + '' + data.description).toLowerCase().indexOf(term.toLowerCase()) !== -1);
  }, [term, data]);

  function handleChange(event: any) {
    setTerm(event.target.value)
  }

  function handleClick() {
    setResult(memoizedResult)
  }

  if (error) {
    return <div>Error: {error}</div>;
  }

  return (
    <>
      <input type="text" id="search" onChange={handleChange} /><button onClick={handleClick}>Search</button>
      {
        result != null ?
          result.map((item: any) => {
            return <ProductBox key={item.id} product={item} />
          })
          :
          ''
      }
    </>
  )
}

export default Search;