package project

type LabelConfig string

func (l LabelConfig) String() string {
	return string(l)
}

var (
	LabelConfigDefaultKeyPointLabels LabelConfig = `
<View>
  <KeyPointLabels name="kp-1" toName="img-1">
    <Label value="Face" background="red" />
    <Label value="Nose" background="green" />
  </KeyPointLabels>
  <Image name="img-1" value="$img" crossOrigin="use-credentials" />
</View>
`

	// NOTE: when you stored your data into your private backend,
	// set env fflag_fix_all_lsdv_4711_cors_errors_accessing_task_data_short=false to disable this feature flag
	// and store credentials in cookies
	// then the data would be fetched back.
	//
	// otherwise, the image tag would be injected with <img crossorigin="annonymous" /> which would not sent any cookies
	// to the backend of hosting your data
	// see the code injected in https://github.com/HumanSignal/label-studio/blob/1.15.0/web/libs/editor/src/components/ImageView/Image.jsx#L82
	LabelConfigDefaultImageClassification LabelConfig = `
<View>
  <Image name="image" value="$image" crossOrigin="use-credentials"/>
  <Choices name="choice" toName="image">
    <Choice value="Adult content"/>
    <Choice value="Weapons" />
    <Choice value="Violence" />
  </Choices>
</View>

`

	LabelConfigDefaultSoundEventDetection LabelConfig = `
<View>
  <Labels name="label" toName="audio" zoom="true" hotkey="ctrl+enter">
    <Label value="Event A" background="red"/>
    <Label value="Event B" background="green"/>
  </Labels>
  <Audio name="audio" value="$audio"/>
</View>
`
)
