import { useEffect } from 'react'
import { useQueryClient } from 'react-query'

function WebSocketClient() {
	const queryClient = useQueryClient()

	useEffect(() => {
		const ws = new WebSocket('ws://localhost:8081/ws')

		ws.onmessage = (event) => {
			const message = JSON.parse(event.data)
			if (message.message === 'ready to update') {
				queryClient.invalidateQueries('randomNumber')
			}
		}

		return () => {
			ws.close()
		}
	}, [queryClient])

	return null
}

export default WebSocketClient
