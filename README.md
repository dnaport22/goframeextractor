#Fisher
Fisher is a extract frames from a video file using ffmpeg library.

The micro-service is a part of Deep Fashion Project and is developed to use along with all other micro-service under this project.

#How to use it
##Prerequisites
 - [ffmpeg](#how-to-install-ffmpeg)

#####NOTE: scroll down to [Appendices](#appendices) for "how tos"

#Setting up local development environment
##Prerequisites
 - [ffmpeg](#how-to-install-ffmpeg)
 - [Golang](#https://golang.org/doc/install)

#####FYI: I used [Goland](https://www.jetbrains.com/go/download/) IDE for developement. I would personally recommend from my experience. However, feel free to pick that works fine for you.
#####NOTE: Scroll down to [Appendices](#appendices) for "how tos".

##Getting the codebase
Now we will clone the repository into our local machine to get started with development.

```bash
> cd INTO/PREFERABLE/DIRECTORY 
> git clone ... fisher
> cd fisher

```

##Running fisher
Let's run pre-compiled fisher code. Call fisher and just pass a video file as a parameter.

```bash
> cd fisher
> bin/fisher -input video.mp4

```
The extracted frames will be stored into a folder with template `VIDEONAME_frames/`.

#Appendices
##How to install ffmpeg
Go to [https://www.ffmpeg.org/download.html](https://www.ffmpeg.org/download.html), follow the instructions on the site to install the library. Any questions, drop a note to me at <navdeep.dhuti@acrotrend.com> or give me a shout if i'm around.