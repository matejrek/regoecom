import useFetch from "../hooks/useFetch";
import ProductBox from "./productBox";
import { Product } from "./productBox";

function AllProducts() {
  const [data, error] = useFetch('products');

  if (error) {
    return <div>Error: {error}</div>;
  }

  return (
    <>
      <div className="row all-products">
        {
          data &&
          data.map((item: Product) => {
            return <ProductBox key={item.id} product={item} />
          })
        }
      </div>
    </>
  )
}

export default AllProducts;