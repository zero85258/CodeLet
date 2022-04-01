import mingus
import mingus.core as core
import mingus.core.progressions as progressions
#from mingus.containers import NoteContainer

#fluidsynth.play_Note(Note("C-5"))
import mingus.midi as midi

print(core.progressions.to_chords(["I", "bIV", "VIIdim7", "II"]))
print(
        core.chords.triad("E", "C")
)

mingus.midi.midi_file_out.write_NoteContainer("tttt.midi",["A","E"])
#MidiFileOut.write_NoteContainer("test.mid", nc)
