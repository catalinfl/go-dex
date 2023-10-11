import React from 'react'

export type DefinitionProps =  {
    word: string,
    typeOfWord: string,
    definition: {
        meaning: string,
        source: string
    }
}

const Definition = (definition: DefinitionProps) => {
  return (
    <div className="bg-orange-400 ronuded-lg p-3 text-white" > 
        <p> <span className="font-bold"> Cuvant: </span> {definition?.word} </p>
        <p> <span className="font-bold"> Parte de vorbire: </span> {definition?.typeOfWord} </p>
        <p> <span className="font-bold"> Inteles: </span> {definition?.definition?.meaning} </p>
        <p> <span className="font-bold"> Surse: </span> {definition?.definition?.source} </p>
    </div>
  )
}

export default Definition