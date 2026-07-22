package mayaBuffers

type ScreenBuffer struct {
    Height int
    Width int
    Font_size int
    MenuBar []string   // Menu bar like [ FILE, EDIT, SETTINGS, HELP]
    RowStart int       // main content will be start from this num. Example line 3 will be starting of main content
    RowEnd int         //  "       "     "  end in this line. Example line 42 will be last shown line
    RowCountBegin int  // This variable represents the line number of a large file content. Like line 23 has something writtem  At start
    RowCountEnd int    //  "       "          "     "   "     "     "      "    "    "        "    "  75  "      "       "      At end. After thisline toolbar at bottom
    InfoBar []string   // Represents absolute last line that shows mode name, row col position and other info
}

type WriterBuffer struct {
    lines []string
}
type FileBuffer struct {
    lines []string
}
