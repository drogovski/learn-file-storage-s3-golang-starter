package main

type Video struct {
	Streams []Stream `json:"streams"`
}

type Stream struct {
	Index              int64            `json:"index"`
	CodecName          string           `json:"codec_name"`
	CodecLongName      string           `json:"codec_long_name"`
	Profile            string           `json:"profile"`
	CodecType          string           `json:"codec_type"`
	CodecTagString     string           `json:"codec_tag_string"`
	CodecTag           string           `json:"codec_tag"`
	Width              *int64           `json:"width,omitempty"`
	Height             *int64           `json:"height,omitempty"`
	CodedWidth         *int64           `json:"coded_width,omitempty"`
	CodedHeight        *int64           `json:"coded_height,omitempty"`
	ClosedCaptions     *int64           `json:"closed_captions,omitempty"`
	FilmGrain          *int64           `json:"film_grain,omitempty"`
	HasBFrames         *int64           `json:"has_b_frames,omitempty"`
	SampleAspectRatio  *string          `json:"sample_aspect_ratio,omitempty"`
	DisplayAspectRatio *string          `json:"display_aspect_ratio,omitempty"`
	PixFmt             *string          `json:"pix_fmt,omitempty"`
	Level              *int64           `json:"level,omitempty"`
	ColorRange         *string          `json:"color_range,omitempty"`
	ColorSpace         *string          `json:"color_space,omitempty"`
	ColorTransfer      *string          `json:"color_transfer,omitempty"`
	ColorPrimaries     *string          `json:"color_primaries,omitempty"`
	ChromaLocation     *string          `json:"chroma_location,omitempty"`
	FieldOrder         *string          `json:"field_order,omitempty"`
	Refs               *int64           `json:"refs,omitempty"`
	IsAVC              *string          `json:"is_avc,omitempty"`
	NalLengthSize      *string          `json:"nal_length_size,omitempty"`
	ID                 string           `json:"id"`
	RFrameRate         string           `json:"r_frame_rate"`
	AvgFrameRate       string           `json:"avg_frame_rate"`
	TimeBase           string           `json:"time_base"`
	StartPts           int64            `json:"start_pts"`
	StartTime          string           `json:"start_time"`
	DurationTs         int64            `json:"duration_ts"`
	Duration           string           `json:"duration"`
	BitRate            string           `json:"bit_rate"`
	BitsPerRawSample   *string          `json:"bits_per_raw_sample,omitempty"`
	NbFrames           string           `json:"nb_frames"`
	ExtradataSize      int64            `json:"extradata_size"`
	Disposition        map[string]int64 `json:"disposition"`
	Tags               Tags             `json:"tags"`
	SampleFmt          *string          `json:"sample_fmt,omitempty"`
	SampleRate         *string          `json:"sample_rate,omitempty"`
	Channels           *int64           `json:"channels,omitempty"`
	ChannelLayout      *string          `json:"channel_layout,omitempty"`
	BitsPerSample      *int64           `json:"bits_per_sample,omitempty"`
	InitialPadding     *int64           `json:"initial_padding,omitempty"`
}

type Tags struct {
	Language    string  `json:"language"`
	HandlerName string  `json:"handler_name"`
	VendorID    string  `json:"vendor_id"`
	Encoder     *string `json:"encoder,omitempty"`
	Timecode    *string `json:"timecode,omitempty"`
}
