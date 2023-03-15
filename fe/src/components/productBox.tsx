import { FC } from "react";
import placeholder from '../../public/placeholder.jpg'

import { useDispatch } from 'react-redux'
import { addproduct } from '../features/cart/cartSlice'


export type Product = {
  id: number;
  title: string;
  description: string;
  slug: string;
  price: number;
  currency: string;
};


export type ProductBoxProps = {
  product: Product
};

export type CartItem = {
  quantity: number,
  product: Product
}

export type CartItemList = CartItem[];


const ProductBox: FC<ProductBoxProps> = (props) => {
  const dispatch = useDispatch()

  function handleAddToCart(item: CartItem) {
    dispatch(addproduct(item))
  }


  return (
    <>
      <div className="card col-md-4">
        <img src={placeholder} className="card-img-top" alt={props.product?.title} />
        <div className="card-body">
          <h5 className="card-title">{props.product?.title}</h5>
          <p className="card-text">{props.product?.description}</p>
          <h3 className="card-text">{props.product?.price} {props.product?.currency}</h3>
          <a href="{props.slug}" className="btn btn-primary">View more</a>
          <button onClick={() => handleAddToCart({ quantity: 1, product: props.product })} >Add to cart</button>
        </div>
      </div>
    </>
  )
}

export default ProductBox;