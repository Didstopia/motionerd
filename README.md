# Motionerd

[![Build Status](https://travis-ci.org/Didstopia/motionerd.svg?branch=master)](https://travis-ci.org/Didstopia/motionerd)
[![codecov](https://codecov.io/gh/Didstopia/motionerd/branch/master/graph/badge.svg)](https://codecov.io/gh/Didstopia/motionerd)

A continuous motion detection system written in GoLang.

**NOTE: _Work in progress._**

## Todo

- [ ] Implement a basic, image-based motion detection system ([link](https://remy.io/blog/how-ive-created-an-intruder-detector-part-2/))
- [ ] Continuously write a video file to the filesystem in chunks
- [ ] Detect motion in real-time, while writing the video file
- [ ] Create a new video file that contains the detected motion, -5/+5 seconds, as well as accounts for a cooldown/offset, in case there's continuous motion going on, but also times out after an X amount of time, so the motion triggers as soon as possible
- [ ] Send email or Pushover notifications with images and/or videos attached of each detected motion "clip", being careful to have a cooldown period for each motion, so it doesn't get out of control and spammy

## Usage

_To be implemented._

## Building

_To be implemented._

## License

See [LICENSE](LICENSE).

## Stargazers

[![Stargazers](https://starcharts.herokuapp.com/Didstopia/motionerd.svg)](https://starcharts.herokuapp.com/Didstopia/motionerd)
