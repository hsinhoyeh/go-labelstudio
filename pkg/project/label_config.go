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
  <Image name="img-1" value="$img" />
</View>
`

	LabelConfigDefaultImageClassification LabelConfig = `
<View>
  <Image name="image" value="$image"/>
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
