import Header from './header/header'
import Footer from './footer/footer'
import { Outlet } from 'react-router-dom'

function Mainlayout() {
  return (
    <>
      <Header />
      <main>
        <div className="container">
          <Outlet />
        </div>
      </main>
      <Footer />
    </>
  );
}

export default Mainlayout;