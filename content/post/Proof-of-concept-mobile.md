---
title: Proof of concept Peernet Android App 
date: 2022-04-13
hero: "/images/peernetmobileheader.png"
excerpt: Proof of concept Peernet Android App 
authors:
  - Akilan Selvacoumar


---

The following blog post is on repo we finally made open source which is a proof of concept 
android app that runs the peernet core internally. The Peernet core team have dedicated 
effort towards implementing the peernet protocol standard in pure go. This effort 
has paid off in making the process of porting Peernet to android a straight forward process.

The peernet android app can do the base funtionality of peernet such as sharing a file 
to the peernet p2p network, utilizing peernets decentralized search to get files from 
peernets p2p network. From a technical sense your android app can escape NAT, has UDT (A great read 
on the benefits of UDT: https://www.cs.uic.edu/~ygu/paper/udt-protocol.pdf) built 
into the app. It's worth looking at the peernet protocol core repo to understand how does 
the peernet protocol work (source: https://github.com/PeernetOfficial/core). 

## How did we link the peernet core to the android application ?
The peernet protcol is in pure Go. This means we can compiler our Go 
libarary of the ARM architecture. To ensure we can interact with 
it from kotlin we need to generate a ```.arr file```. To do this 
we used gomobile to generate the appropirate files and then we just 
linked them to the android project. The MainActivity calls the 
go main function to start the network and create a local server 
on the phone (ex: 127.0.0.1:5125). Once the server is running 
koltin interacts with the peernet local server via REST API calls
(Peernet rest API calls docs: https://github.com/PeernetOfficial/core/tree/master/webapi#readme).

## User perpective 
The android apk released is not the best design nor has functinality. This is 
as we mentioned before it's a proof of concept. The app can currently 
list the numbers of peers the phone can detect in the peernet public network. 
A user can add a file to the [warehouse](https://github.com/PeernetOfficial/core/tree/master/warehouse#readme) and [blockchain](https://github.com/PeernetOfficial/core/tree/master/blockchain#readme). The user can view recently added file to the network/search 
for files in the network using file keywords (at the currect release filename. The following [README.md](https://github.com/PeernetOfficial/core/tree/master/search#readme) gives a better idea 
on how we do search indexing and normalization). 

## Project information 
- Source code: https://github.com/PeernetOfficial/core-android
- Released APK file: https://github.com/PeernetOfficial/core-android/releases/tag/1.0.0-beta

### Note: Join the peernet discord channel to discuss more: 
- https://discord.gg/nEzUm2gADh

### We hope you can join us to improve peernet to become a simple and effective p2p network :) 

Landing page: https://peernet.org
---

{{< css.inline >}}
<style>
.canon { background: white; width: 100%; height: auto}

</style>
{{< /css.inline >}}
=