import React, { Suspense } from 'react';
import ReactDOM from 'react-dom/client'
import Search from '../components/search';
import Popupcart from '../components/popupcart';


const AllProducts = React.lazy(() => import('../components/allProducts'));


function Home() {
  return (
    <>
      <div className="row">
        <div className="col-md-12">
          <Popupcart />
        </div>
      </div>
      <div className="row">
        <div className="col-md-12">
          <Search />
        </div>
      </div>
      <div className="row">
        <div className="col-md-12">
          <h2>All products</h2>
          <Suspense fallback={<div>Loading...</div>}>
            <AllProducts />
          </Suspense>
        </div>
      </div>
    </>
  );
}

export default Home;