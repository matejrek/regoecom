import Navigation from '../../components/navigation.tsx'


export default function Header() {
  return (
    <>
      <div className="container">
        <h1><a href="/">E-commerce</a></h1>
        <Navigation />
      </div>
    </>
  );
}