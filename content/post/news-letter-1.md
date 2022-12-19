---
title: Peernet Newsletter (1)
date: 2021-05-05
hero: "/images/peernetheader.png"
excerpt: Peernet Newsletter (1)
authors:
  - Peter Kleissner

---

Welcome to the first Peernet newsletter! We have published the pitch deck here: https://peernet.org/dl/Peernet%20Deck.pdf 🚀

## ICO
We are preparing the ICO Seed Round which is scheduled to close end of Q2 2021. There is currently 200k EUR of 400k EUR pre-committed. If you are interested, please contact us! Note that minimum participation is 50k. The funds will be used to finance the first version of the Peernet Browser. 

## Upcoming Alpha 2
The upcoming milestone alpha 2 brings improved connectivity. We are adding UPnP support (for now IPv4 only) and UDP hole punching. Down the road we’ll add support for IPv6 UPnP, NAT-PMP (= IPv4 only) and PCP (= successor of NAT-PMP), which are different protocols that allow to forward ports.

Forwarding ports is important for improved connectivity in case the computer is behind a NAT, which is typically the case for home networks. UDP hole punching should do the trick most of the times, but will fail on symmetric NATs.

## Fundamental Question: Server/client mode?
Kazaa has supernodes, IPFS has a server/client mode.

Peernet will not. Every peer is equal. There is not good reason to elevate some peers to a special status. If a peer is not accessible, it is not accessible. It is a dangerous road and sets bad precedent to artificially segregate peers into different categories based on their temporary accessibility. If certain data is only shared by “supernodes” which typically are only a small percentage of all participants in the network, it would significantly lower the bar for a variety of attacks (including disruption, isolation, spam).

Our approach is to instead improve connectivity for all. The Peernet protocol is designed to be small and efficient - a single packet establishes a connection! In the future we will add a special Proxy message to the protocol that improves connectivity in heavily firewalled environments.

## What to expect in the near future
We plan to have the first early version of the Peernet Browser ready by the summer! Until then you can use the command line tool to play around. There is not yet much useful functionality since we are working on the very basic problems, but the Peernet network is already live.

In the next few weeks we’ll start working on the higher level functionality such as publishing files via the user’s own blockchain, distributing the blockchain, providing discovery of data, exchange of files etc. The best way to think about the development of Peernet is in layers. We solved the initial protocol, the encryption, peer identity (public/private key), basic connectivity and soon we will move on to the DHT (distributed hash table).

## How you can help
Spread the word about Peernet! Participate in the forum (link below) and the technical discussions! If you are a developer and you want to help please reach out! If you have a spare server, why don’t you run a peer?

## Our presence
Here are links to our legacy internet presence:
- Forum: talk.peernet.org
- Blog: blog.peernet.org
- GitHub: github.com/PeernetOfficial
- Twitter: @PeernetOfficial

## Current statistics:
- 6 Peers
- 48 Newsletter subscribers
- 83 Twitter followers