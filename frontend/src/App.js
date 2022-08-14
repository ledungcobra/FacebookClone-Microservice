function App() {
  const get = async () => {
    const result = await fetch("http://localhost:8080/api");
    console.log(result);
  };
  get();
  return <div>welcome to frontend asd</div>;
}

export default App;
