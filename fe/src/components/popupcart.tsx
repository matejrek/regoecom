import { useState } from 'react';
import type { RootState } from '../store'
import { useSelector, useDispatch } from 'react-redux'
import { removeproduct, increment, decrement } from '../features/cart/cartSlice'

export type CartItem = {
  quantity: number,
  product: Product
}

export type CartItemList = CartItem[];


function Popupcart() {
  const [cart, setCart] = useState<CartItemList | null>(null);
  const dispatch = useDispatch()
  const shoppingCart = useSelector((state: RootState) => state.cart.data)

  function handleRemoveFromCart(id: number) {
    dispatch(removeproduct(id))
  }

  function handleIncrementItemInCart(item: CartItem) {
    dispatch(increment(item))
  }

  function handleDecrementItemInCart(item: CartItem) {
    dispatch(decrement(item))
  }

  return (
    <>
      <ul>
        {
          shoppingCart != null ?
            shoppingCart.map((item) => {
              return (
                <li key={"li-" + item.product.id}>{item.product.title} ({item.quantity}) --- ({item.product.price ? item.quantity * item.product.price : ''} {item.product.currency}) ---
                  <button onClick={() => { handleDecrementItemInCart(item) }}>-</button>
                  <button onClick={() => { handleIncrementItemInCart(item) }}>+</button>
                  <button onClick={() => { handleRemoveFromCart(item.product.id) }}>x</button>
                </li>
              )
            })
            :
            ""
        }
      </ul>

      TOTAL: --{/*
        shoppingCart != null ?
          shoppingCart.map((item, sum: number = 0) => {
            sum = sum + (Number(item.product.price) * Number(item.quantity))
            return sum
          })

          :
          ''
      */}
      EUR
    </>
  )
}

export default Popupcart;