package main

var news []News

func main() {
	news = append(news, getNewYorkTimes().Results...)

	// for key, pieceOfNews := range reponse.Results {
	// 	news = append(news, pieceOfNews)
	// 	key := 1
	// }

	handleRequests()
}
