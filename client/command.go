package client

import "github.com/desertbit/grumble"

var tlsClient *TLSClient

const address = "121.15.171.87:12345"

func init() {
	tlsClient = NewTLSClient("", "", "")
}

func CreateUploadCommand() *grumble.Command {
	fc := new(grumble.Command)
	fc.Name = "FileUpload"
	fc.Help = "Upload a file to the system"
	fc.Args = func(a *grumble.Args) {
		a.String("fileName", "File's Name")
	}
	fc.Flags = func(f *grumble.Flags) {
		f.Uint64("p", "threads", 1, "Threads used in uploading files")
	}
	fc.Run = func(c *grumble.Context) error {
		conn, err := tlsClient.Dial(address)
		if err != nil {
			return err
		}

		uploadTLSConnQueue <- &TLSConn{
			Conn:    conn,
			Context: c,
		}
		return nil
	}
	return fc
}

func CreateListCommand() *grumble.Command {
	fc := new(grumble.Command)
	fc.Name = "FileList"
	fc.Help = "Get File list"
	fc.Run = func(c *grumble.Context) error {
		conn, err := tlsClient.Dial(address)
		if err != nil {
			return err
		}

		listTLSConnQueue <- &TLSConn{
			Conn:    conn,
			Context: c,
		}
		return nil
	}

	return fc
}

func CreateDeleteCommand() *grumble.Command {
	fc := new(grumble.Command)
	fc.Name = "FileDel"
	fc.Help = "Delete a File"
	fc.Args = func(a *grumble.Args) {
		a.Uint64("fileId", "File's Id")
	}
	fc.Run = func(c *grumble.Context) error {
		conn, err := tlsClient.Dial(address)
		if err != nil {
			return err
		}

		deleteTLSConnQueue <- &TLSConn{
			Conn:    conn,
			Context: c,
		}
		return nil
	}

	return fc
}

func CreateDownloadCommand() *grumble.Command {
	fc := new(grumble.Command)
	fc.Name = "FileUpload"
	fc.Help = "Download a file from the system"
	fc.Args = func(a *grumble.Args) {
		a.Uint64("fileId", "File's Id")
		a.String("fileLocation", "File's Destination")
	}
	fc.Flags = func(f *grumble.Flags) {
		f.Uint64("p", "threads", 1, "Threads used in uploading files")
	}
	fc.Run = func(c *grumble.Context) error {
		conn, err := tlsClient.Dial(address)
		if err != nil {
			return err
		}

		downloadTLSConnQueue <- &TLSConn{
			Conn:    conn,
			Context: c,
		}
		return nil
	}
	return fc
}
