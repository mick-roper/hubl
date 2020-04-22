module github.com/mick-roper/hubl

go 1.14

require github.com/mick-roper/hubl/pkg/common v1.0.0
require github.com/mick-roper/hubl/pkg/web v1.0.0

replace github.com/mick-roper/hubl/pkg/common => ./pkg/common
replace github.com/mick-roper/hubl/pkg/web => ./pkg/web