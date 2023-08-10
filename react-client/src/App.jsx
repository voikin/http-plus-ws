import { useQuery, QueryClient, QueryClientProvider } from 'react-query'
import WebSocketClient from './WebSocketClient'
import './App.css'
import { useState } from 'react'

const queryClient = new QueryClient()

async function fetchRandomNumber() {
	const response = await fetch('http://localhost:8080/random')
	console.log(response)
	const data = await response.json()
	return data.randomNumber
}

function App() {
	const [number, setNumber] = useState(0)
	const { isLoading } = useQuery('randomNumber', fetchRandomNumber, {
		onSuccess: (num) => setNumber(num),
	})

	return (
		<div className='App'>
			{isLoading ? (
				<p>loading...</p>
			) : (
				<progress value={number} max='10'></progress>
			)}
			<WebSocketClient />
		</div>
	)
}

function AppWithQuery() {
	return (
		<QueryClientProvider client={queryClient}>
			<App />
		</QueryClientProvider>
	)
}

export default AppWithQuery
