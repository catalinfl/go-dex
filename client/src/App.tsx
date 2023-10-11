import { FormEvent, useEffect, useState } from "react"
import Definition, { DefinitionProps } from "./components/Definition"
import axios from "axios"

function App() {

  const [word, setWord] = useState<string>("")
  const [information, setInformation] = useState<DefinitionProps[]>([])
  console.log(word)
  const handleSubmit = (e: FormEvent<HTMLButtonElement>) => {
    e.preventDefault()
    if (word.length > 3) {
      axios.get(`http://localhost:3000/search/${word}`)
                .then((res) => setInformation(res.data))
                .catch((err) => console.log(err))    
    
    }
    }
  
  useEffect(() => {
    console.log(information)
  }, [information])
  
  console.log(information)

  return (
    <>
      <div className="flex flex-col max-w-2xl mx-auto p-4 sm:p-0"> 
      <div className="max-w-2xl flex flex-col mx-auto text-center mt-4 ">
        <p className="text-3xl font-bold text-orange-300 text-transparent bg-clip-text bg-gradient-to-r from-red-200 to-orange-500"> definition.ro </p>
        <p className="text-[1.5rem] font-bold mt-4 text-orange-300"> Gasește o definiție pentru orice cuvânt dorești în timp real </p>
      </div>
      <div className="bg-orange-400 flex flex-col gap-3 p-3 mt-4">
        <p className="text-white cursor-pointer ml-2 mt-3 text-center"> Caută un cuvant</p>
        <div className="flex justify-center"> 
        <input type="text" onChange={(e: React.ChangeEvent<HTMLInputElement>) => setWord(e.target.value)} className="flex w-[75%] outline-none  rounded-lg px-2 py-1" />
        </div>
        <p className="text-white cursor-pointer mt-1"> </p>
        <button type="button" className="bg-orange-300 mx-auto p-3 text-white" onClick={(e: FormEvent<HTMLButtonElement>) => handleSubmit(e)}> Caută </button>
      </div>
      <div className="flex flex-col gap-2 mt-4">
        {information?.map((information: DefinitionProps) => <Definition key={information.definition.meaning} {...information} />)}
      </div>
      </div>
    </>
  )
}

export default App


