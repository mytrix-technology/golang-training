package main

//type UserMessage struct {
//	Location   string
//	WaveHeight int
//}
//
//func makeHost(port int, randomness io.Reader) (host.Host, error) {
//	// Creates a new RSA key pair for this host.
//	prvKey, _, err := crypto.GenerateKeyPairWithReader(crypto.RSA, 2048, randomness)
//	if err != nil {
//		log.Println(err)
//		return nil, err
//	}
//
//	// 0.0.0.0 will listen on any interface device.
//	sourceMultiAddr, _ := multiaddr.NewMultiaddr(fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", port))
//
//	// libp2p.New constructs a new libp2p Host.
//	// Other options can be added here.
//	return libp2p.New(
//		libp2p.ListenAddrs(sourceMultiAddr),
//		libp2p.Identity(prvKey),
//	)
//}
//
//func handleStream(s network.Stream) {
//	log.Println("Stream detected")
//
//	// Create a buffer stream for non-blocking read and write.
//	rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))
//
//	go readData(rw)
//	go writeData(rw)
//
//	// stream 's' will stay open until you close it (or the other side closes it).
//}
//
//func readData(rw *bufio.ReadWriter) {
//	for {
//		str, err := rw.ReadString('\n')
//
//		// If the channel is closed or we get an EOF, return
//		if err == io.EOF {
//			return
//		}
//
//		if str == "" {
//			continue
//		}
//
//		if str != "\n" {
//			var chain Blockchain
//			if err := json.Unmarshal([]byte(str), &chain); err != nil {
//				log.Println(err)
//				continue
//			}
//
//			mutex.Lock()
//			if len(chain.Chain) > len(mychain.Chain) {
//				mychain = chain
//				pretty.Println(mychain)
//			}
//			mutex.Unlock()
//		}
//
//	}
//}
//
//func writeData(rw *bufio.ReadWriter) {
//
//	stdReader := bufio.NewReader(os.Stdin)
//
//	for {
//		fmt.Print("> ")
//		sendData, err := stdReader.ReadString('\n')
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		sendData = strings.Replace(sendData, "\n", "", -1)
//		var userMsg UserMessage
//
//		if err := json.Unmarshal([]byte(sendData), &userMsg); err != nil {
//			log.Println(err)
//			continue
//		}
//
//		mutex.Lock()
//		mychain.appendBlock(userMsg.Location, userMsg.WaveHeight)
//		if !mychain.isValid() {
//			pretty.Println(mychain)
//			log.Println("Chain isn't valid anymore! Help meee!")
//			return
//		}
//		mutex.Unlock()
//
//		bytes, err := json.Marshal(mychain)
//		log.Println(string(bytes))
//		if err != nil {
//			log.Println(err)
//		}
//
//		pretty.Println(mychain)
//
//		mutex.Lock()
//		rw.WriteString(fmt.Sprintf("%s\n", string(bytes)))
//		rw.Flush()
//		mutex.Unlock()
//	}
//}
//
//func startPeer(ctx context.Context, h host.Host, streamHandler network.StreamHandler) {
//	// Set a function as stream handler.
//	// This function is called when a peer connects, and starts a stream with this protocol.
//	// Only applies on the receiving side.
//	h.SetStreamHandler("/chat/1.0.0", streamHandler)
//
//	// Let's get the actual TCP port from our listen multiaddr, in case we're using 0 (default; random available port).
//	var port string
//	for _, la := range h.Network().ListenAddresses() {
//		if p, err := la.ValueForProtocol(multiaddr.P_TCP); err == nil {
//			port = p
//			break
//		}
//	}
//
//	if port == "" {
//		log.Println("was not able to find actual local port")
//		return
//	}
//
//	log.Printf("Run './simple-blockchain -d /ip4/127.0.0.1/tcp/%v/p2p/%s' on another console.\n", port, h.ID())
//	log.Println("You can replace 127.0.0.1 with public IP as well.")
//	log.Println("Waiting for incoming connection")
//	log.Println()
//}
//
//func startPeerAndConnect(ctx context.Context, h host.Host, destination string) (*bufio.ReadWriter, error) {
//	log.Println("This node's multiaddresses:")
//	for _, la := range h.Addrs() {
//		log.Printf(" - %v\n", la)
//	}
//	log.Println()
//
//	// Turn the destination into a multiaddr.
//	maddr, err := multiaddr.NewMultiaddr(destination)
//	if err != nil {
//		log.Println(err)
//		return nil, err
//	}
//
//	// Extract the peer ID from the multiaddr.
//	info, err := peer.AddrInfoFromP2pAddr(maddr)
//	if err != nil {
//		log.Println(err)
//		return nil, err
//	}
//
//	// Add the destination's peer multiaddress in the peerstore.
//	// This will be used during connection and stream creation by libp2p.
//	h.Peerstore().AddAddrs(info.ID, info.Addrs, peerstore.PermanentAddrTTL)
//
//	// Start a stream with the destination.
//	// Multiaddress of the destination peer is fetched from the peerstore using 'peerId'.
//	s, err := h.NewStream(context.Background(), info.ID, "/chat/1.0.0")
//	if err != nil {
//		log.Println(err)
//		return nil, err
//	}
//	log.Println("Established connection to destination")
//
//	// Create a buffered stream so that read and writes are non-blocking.
//	rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))
//
//	return rw, nil
//}
