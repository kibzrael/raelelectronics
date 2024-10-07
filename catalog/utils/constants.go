package utils

type Color struct {
	Name string
	Hex  string
}

var DisplayTypes = []string{
	"LED IPS",
	"OLED",
	"LED TN",
}

var DisplaySizes = []uint64{
	12,
	13,
	14,
	15,
	16,
}

var Colors = []Color{
	{Name: "Black", Hex: "#000000"},
	{Name: "Silver", Hex: "#000000"},
	{Name: "White", Hex: "#FFFFFF"},
	{Name: "Gray", Hex: "#000000"},
	{Name: "Gold", Hex: "#000000"},
	{Name: "Blue", Hex: "#000000"},
	{Name: "Orange", Hex: "#000000"},
}

var GPUBrands = []string{
	"Intel",
	"Nvidia",
	"Amd",
	"Apple",
	"Arm",
}

var CPUs = []string{
	"Intel i3",
	"Intel i5",
	"Intel i7",
	"Amd Ryzen 3",
	"Amd Ryzen 5",
	"Amd Ryzen 7",
	"Amd Ryzen 9",
	"Arm-based Apple",
	"Arm-based Snapdragon",
	"Intel Ultra 5",
	"Intel Ultra 7",
	"Intel Ultra 9",
}

var BRANDS []string = []string{"Lenovo", "HP", "Dell", "Apple", "Samsung", "Asus", "MSI", "Acer", "Microsoft", "Alienware", "Huawei", "LG", "Gigabyte", "Razer", "Framework", "Fujitsu", "Xiaomi", "Google"}

var MemorySizes = []uint64{
	8,
	12,
	16,
	32,
}
