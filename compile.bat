@echo on
echo Compiling ProPresenter 7 protobuf definitions into Go code...
protoc --proto_path=../ProPresenter7-Proto/proto/ --go_out=./ action.proto alignmentGuide.proto audio.proto background.proto baseDocument.proto basicTypes.proto calendar.proto ccli.proto cue.proto digitalAudio.proto effects.proto graphicsData.proto groups.proto hotKey.proto input.proto labels.proto layers.proto library.proto liveVideoPlaylist.proto masks.proto messages.proto planningCenter.proto playlist.proto presentation.proto presentationSlide.proto proAudienceLook.proto proCore.proto proMask.proto propresenter.proto propSlide.proto proscreen.proto proworkspace.proto recording.proto screens.proto slide.proto stage.proto targets.proto template.proto testPattern.proto timers.proto timestamp.proto workspace.proto
ren Pro7Protobuf proto
pause