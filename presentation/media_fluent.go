package presentation

import (
	"github.com/connerohnesorge/goffice/presentation/elements"
	"github.com/connerohnesorge/goffice/presentation/parts"
)

// Video represents a video shape on a slide with a fluent API.

type Video struct {
	slide *parts.SlidePart

	part parts.MediaPart

	pic *elements.Picture

	props elements.MediaProperties
}

// AddVideo adds a video to the specified slide.

func (d *Document) AddVideo(slideIndex int, filePath string) (*Video, error) {

	part, err := d.AddVideoFromFile(slideIndex, filePath, nil)

	if err != nil {

		return nil, err

	}

	slide, err := d.GetSlide(slideIndex)

	if err != nil {

		return nil, err

	}

	// Add the video to the slide's XML

	pic := slide.AddVideo(part.RelationshipID(), "")

	v := &Video{

		slide: slide,

		part: part,

		pic: pic,

		props: elements.MediaProperties{

			EmbedRelId: part.RelationshipID(),
		},
	}

	return v, nil

}

// SetPosition sets the position of the video in EMUs.

func (v *Video) SetPosition(x, y int) *Video {

	v.pic.SetPosition(x, y)

	return v

}

// SetSize sets the size of the video in EMUs.

func (v *Video) SetSize(w, h int) *Video {

	v.pic.SetSize(w, h)

	return v

}

// SetPoster sets the poster image for the video.

func (v *Video) SetPoster(img *parts.ImagePart) *Video {

	v.pic.SetRelId(img.RelationshipID())

	return v

}

// SetAutoStart sets whether the video starts automatically.

func (v *Video) SetAutoStart(auto bool) *Video {

	v.props.AutoStart = auto

	v.applyProps()

	return v

}

// SetLoop sets whether the video loops.

func (v *Video) SetLoop(loop bool) *Video {

	v.props.Loop = loop

	v.applyProps()

	return v

}

// SetMuted sets whether the video is muted.

func (v *Video) SetMuted(muted bool) *Video {

	v.props.Muted = muted

	v.applyProps()

	return v

}

// SetVolume sets the video volume (0-100000).

func (v *Video) SetVolume(vol int) *Video {

	v.props.Volume = vol

	v.applyProps()

	return v

}

func (v *Video) applyProps() {

	anvp := v.pic.NonVisualPictureProperties().ApplicationNonVisualProperties()

	anvp.SetMediaProperties(v.props)

}

// Audio represents an audio shape on a slide with a fluent API.

type Audio struct {
	slide *parts.SlidePart

	part parts.MediaPart

	pic *elements.Picture

	props elements.MediaProperties
}

// AddAudio adds an audio to the specified slide.

func (d *Document) AddAudio(slideIndex int, filePath string) (*Audio, error) {

	part, err := d.AddAudioFromFile(slideIndex, filePath, nil)

	if err != nil {

		return nil, err

	}

	slide, err := d.GetSlide(slideIndex)

	if err != nil {

		return nil, err

	}

	// Add the audio to the slide's XML

	pic := slide.AddAudio(part.RelationshipID(), "")

	a := &Audio{

		slide: slide,

		part: part,

		pic: pic,

		props: elements.MediaProperties{

			EmbedRelId: part.RelationshipID(),
		},
	}

	return a, nil

}

// SetPosition sets the position of the audio icon in EMUs.

func (a *Audio) SetPosition(x, y int) *Audio {

	a.pic.SetPosition(x, y)

	return a

}

// SetSize sets the size of the audio icon in EMUs.

func (a *Audio) SetSize(w, h int) *Audio {

	a.pic.SetSize(w, h)

	return a

}

// SetLoop sets whether the audio loops.

func (a *Audio) SetLoop(loop bool) *Audio {

	a.props.Loop = loop

	a.applyProps()

	return a

}

// SetMuted sets whether the audio is muted.

func (a *Audio) SetMuted(muted bool) *Audio {

	a.props.Muted = muted

	a.applyProps()

	return a

}

// SetVolume sets the audio volume (0-100000).

func (a *Audio) SetVolume(vol int) *Audio {

	a.props.Volume = vol

	a.applyProps()

	return a

}

func (a *Audio) applyProps() {

	anvp := a.pic.NonVisualPictureProperties().ApplicationNonVisualProperties()

	anvp.SetMediaProperties(a.props)

}
