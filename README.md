<h1>Fisher</h1>
Fisher is a extract frames from a video file using ffmpeg library.

The micro-service is a part of Deep Fashion Project and is developed to use along with all other micro-service under this project.

<h1>How to use it</h1>
<h2>Prerequisites</h2>
 - [ffmpeg](#how-to-install-ffmpeg)

<p>NOTE: scroll down to [Appendices](#appendices) for "how tos"</p>

<h1>Setting up local development environment</h1>
<h2>Prerequisites</h2>
 - [ffmpeg](#how-to-install-ffmpeg)
 - [Golang](#https://golang.org/doc/install)

<b>FYI: I used [Goland](https://www.jetbrains.com/go/download/) IDE for developement. I would personally recommend from my experience. However, feel free to pick that works fine for you.</b>
<b>NOTE: Scroll down to [Appendices](#appendices) for "how tos".</b>

<h2>Getting the codebase</h2>
Now we will clone the repository into our local machine to get started with development.

```bash
> cd INTO/PREFERABLE/DIRECTORY 
> git clone ... fisher
> cd fisher

```

<h2>Running fisher</h2>
Let's run pre-compiled fisher code. Call fisher and just pass a video file as a parameter.

```bash
> cd fisher
> bin/fisher -input video.mp4

```
The extracted frames will be stored into a folder with template `VIDEONAME_frames/`.

<h1>Appendices</h1>
<h2>How to install ffmpeg</h2>
Go to [https://www.ffmpeg.org/download.html](https://www.ffmpeg.org/download.html), follow the instructions on the site to install the library. Any questions, drop a note to me at <navdeep.dhuti@acrotrend.com> or give me a shout if i'm around.
