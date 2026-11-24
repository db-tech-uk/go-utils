package fileutils

func ByteSizeConvert(fileBytes int64, units string) float64 {
	var oneMb float64 = 1024

	bytes := float64(fileBytes)

	kilobytes := bytes / oneMb

	megabytes := kilobytes / oneMb

	gigabytes := megabytes / oneMb

	terabytes := gigabytes / oneMb

	petabytes := terabytes / oneMb

	exabytes := petabytes / oneMb

	switch units {
	case "kb":
		return bytes / oneMb

	case "mb":
		return kilobytes / oneMb

	case "gb":
		return megabytes / oneMb

	case "tb":
		return gigabytes / oneMb

	case "pb":
		return terabytes / oneMb

	case "xb":
		return petabytes / oneMb

	case "zb":
		return exabytes / oneMb
	}

	return bytes
}
