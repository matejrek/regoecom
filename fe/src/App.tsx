import './App.css'
import Mainlayout from './layout/mainLayout'

import Categories from './pages/categories'
import {
  createBrowserRouter,
  RouterProvider,
  Route,
  Link,
} from "react-router-dom";
import Home from './pages/home'
import Products from './pages/products'
import Cart from './pages/cart'
import Register from './pages/auth/register';
import Login from './pages/auth/login';
import Admin from './pages/admin/admin';
import ProtectedRoute from './pages/admin/private-route';

const router = createBrowserRouter([
  {
    path: "/",
    element: <Mainlayout />,
    children: [
      {
        index: true,
        element: <Home />,
      },
      {
        path: "/categories",
        element: <Categories />,
      },
      {
        path: "/products",
        element: <Products />,
      },
      {
        path: "/cart",
        element: <Cart />,
      },
      {
        path: "/register",
        element: <Register />,
      },
      {
        path: "/login",
        element: <Login />
      },
      {
        path: "/admin",
        element: <ProtectedRoute element={<Admin />} /> /*<Admin />*/
      },
    ],
  },
]);

function App() {

  return (
    <div className="App">
      <RouterProvider router={router} />
    </div >
  )
}

export default App
