package liveMedia

type MPEGVideoStreamParser struct {
	StreamParser
	buffTo            []byte
	startOfFrame      []byte
	numTruncatedBytes uint
	usingSource       *H264VideoStreamFramer
}

func (this *MPEGVideoStreamParser) registerReadInterest(to []byte, maxSize uint) {
	//this.StartOfFrame = this.buffTo
	//this.SavedTo = to
	//fLimit = to + maxSize
}

func (this *MPEGVideoStreamParser) NumTruncatedBytes() uint {
	return this.numTruncatedBytes
}

func (this *MPEGVideoStreamParser) saveByte(ubyte uint) {
}

func (this *MPEGVideoStreamParser) save4Bytes(word uint) {
}

func (this *MPEGVideoStreamParser) curFrameSize() uint {
	return 0
}

func (this *MPEGVideoStreamParser) setParseState() {
	this.saveParserState()
}
