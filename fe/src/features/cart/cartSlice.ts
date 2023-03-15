import { createSlice } from '@reduxjs/toolkit'
import type { PayloadAction } from '@reduxjs/toolkit'
import { Product } from '../../hooks/________types'
import { CartItem } from '../../hooks/________types'


export interface CartState {
  data: Array<CartItem>
}

const initialState: CartState = {
  data: [],
}

/*
A "slice" is a collection of Redux reducer logic and actions for a single feature in your app, 
typically defined together in a single file. The name comes from splitting up the root Redux 
state object into multiple "slices" of state.


PURE FUNCTIONS: Its return value is only determined by its input values
*/

export const cartSlice = createSlice({
  name: 'cart',
  initialState,
  reducers: {
    addproduct: (state, action: PayloadAction<CartItem>) => {
      //state.data.push(action.payload)
      //check for id, if it doesnt exist, add, if it does increase quantity
      let addNew = true;
      state.data.map((item) => {
        if (item.product.id == action.payload.product.id) {
          item.quantity++;
          addNew = false;
        }
      })
      if (addNew) {
        state.data.push(action.payload)
      }
    },
    increment: (state, action: PayloadAction<CartItem>) => {
      state.data.map((item) => {
        if (item.product.id == action.payload.product.id) {
          item.quantity++;
        }
      })
    },
    decrement: (state, action: PayloadAction<CartItem>) => {
      state.data.map((item) => {
        if (item.product.id == action.payload.product.id) {
          item.quantity--;
        }
      })
    },
    removeproduct: (state, action: PayloadAction<Number>) => {
      //find product id and remove it
      state.data = state.data.filter(item => item.product.id != action.payload)
    }
  }
})

export const { addproduct, removeproduct, increment, decrement } = cartSlice.actions

export default cartSlice.reducer