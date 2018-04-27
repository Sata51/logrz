package logrz

import "testing"

func TestLevel_String(t *testing.T) {
	tests := []struct {
		name  string
		level Level
		want  string
	}{
		{
			name:  "DebugLevel",
			level: DebugLevel,
			want:  "DEBUG",
		},
		{
			name:  "InfoLevel",
			level: InfoLevel,
			want:  "INFO",
		},
		{
			name:  "WarnLevel",
			level: WarnLevel,
			want:  "WARN",
		},
		{
			name:  "ErrorLevel",
			level: ErrorLevel,
			want:  "ERROR",
		},
		{
			name:  "FatalLevel",
			level: FatalLevel,
			want:  "FATAL",
		},
		{
			name:  "PanicLevel",
			level: PanicLevel,
			want:  "PANIC",
		},
		{
			name:  "TraceLevel",
			level: TraceLevel,
			want:  "TRACE",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.level.String(); got != tt.want {
				t.Errorf("Level.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseLevel(t *testing.T) {
	type args struct {
		lvl string
	}
	tests := []struct {
		name    string
		args    args
		want    Level
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "PANIC",
			args:    args{lvl: "PANIC"},
			want:    PanicLevel,
			wantErr: false,
		},
		{
			name:    "FATAL",
			args:    args{lvl: "FATAL"},
			want:    FatalLevel,
			wantErr: false,
		},
		{
			name:    "ERROR",
			args:    args{lvl: "ERROR"},
			want:    ErrorLevel,
			wantErr: false,
		},
		{
			name:    "WARN",
			args:    args{lvl: "WARN"},
			want:    WarnLevel,
			wantErr: false,
		},
		{
			name:    "WARNING",
			args:    args{lvl: "WARNING"},
			want:    WarnLevel,
			wantErr: false,
		},
		{
			name:    "INFO",
			args:    args{lvl: "INFO"},
			want:    InfoLevel,
			wantErr: false,
		},
		{
			name:    "DEBUG",
			args:    args{lvl: "DEBUG"},
			want:    DebugLevel,
			wantErr: false,
		},
		{
			name:    "TRACE",
			args:    args{lvl: "TRACE"},
			want:    TraceLevel,
			wantErr: false,
		},
		{
			name:    "ERR",
			args:    args{lvl: "ERR"},
			want:    PanicLevel,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseLevel(tt.args.lvl)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseLevel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}
