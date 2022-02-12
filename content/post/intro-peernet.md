---
title: Welcome to Peernet 
date: 2022-02-12
hero: "/images/peernetheader.png"
excerpt: Introduction of peernet the community
authors:
  - Akilan Selvacoumar

---

This is our first Peernet blog post. Peernet is a p2p network with no intermediaries. Our main focus is to have a browser tailor made for p2p networks. The goal of Peernet is to have a user-friendly experience while sharing content and being completely decentralized.  Peernet is founded by [Peter Kleissner](https://peterkleissner.com/about/) who is also the founder of [Intelligence X](https://intelx.io/about). The p2p space is indeed crowded by multiple breakthroughs and implementations. We are working towards being user friendly, building a custom open standard protocol based on other implementations and breakthroughs in the p2p and file sharing space. 

We currently have our own p2p standard built from the ground up and is completely open source. Our p2p implementation is released under the MIT license. The following is the Peernet protocol and Command line implementation we officially support: 
- Core: The Go implementation of the Peernet Open standard. (https://github.com/PeernetOfficial/core)
- Cmd: The command line implementation of the Core repository.  
(https://github.com/PeernetOfficial/Cmd)

To have a great understanding of Peernets protocol design philosophy the whitepaper we provide gives a good understanding of inner understanding of the protocol.
- Whitepaper: https://peernet.org/dl/Peernet%20Whitepaper.pdf

The Peernet browser is designed to ensure that any user shares content and views content on the Peernet protocol. We found that using standard Web2 browsers will not do the trick to transition to a decentralized p2p (i.e Web3). Due to this we have dedicated our efforts to build a browser from the ground up to ensure users can share files with ease. The Peernet browser currently only supports Windows (This is due to the time factor it takes to design a browser from scratch and with native performance). 

There are a few things we are glad we did. It is currently possible to run the Peernet core implementation on most Architecture. This is because the entire Peernet implementation is in pure Go (Including the key-value storage used). We use classic techniques such [UDP holepunching](https://en.wikipedia.org/wiki/UDP_hole_punching) and [UPNP](https://en.wikipedia.org/wiki/Universal_Plug_and_Play) to ensure any node behind NAT can communicate with other nodes in the network. The Peernet Core also currently runs Natively on Android according to internal tests we have done. Another feature we hope to be extremely useful is the decentralized search in a p2p network. 

There are more blogs to follow on topics such as future plans, inner explanation of certain concepts we use in our p2p network etc…

If you are interested to join Peernet as developer or interested to discuss the inner working of the Peernet protocol. 
Join our [discord channel](https://discord.gg/2hswTarK)
(We are excited to welcome you to our community :))

We are also currently looking for investors. If interested email: info@peernet.org 


