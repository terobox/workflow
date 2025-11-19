
function App() {
  const apiUrl = import.meta.env.VITE_API_URL

  return (
    <>
      <div className="text-xl text-muted-foreground text-center mt-10">welcome to workflow web.</div>
      <div className="text-xl text-muted-foreground text-center mt-10">
        {apiUrl || "error: none"}
      </div>
    </>
  )
}

export default App
