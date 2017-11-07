// Copyright 2017 syzkaller project authors. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.

// I heard you like shell...

//go:generate bash -c "echo -e '// AUTOGENERATED FROM executor/common_linux.h\npackage csource\nvar commonHeaderLinux = `' > linux_common.go; cat ../../executor/common_linux.h | sed -e '/#include \"common.h\"/ {' -e 'r ../../executor/common.h' -e 'd' -e '}' | sed -e '/#include \"common_kvm_amd64.h\"/ {' -e 'r ../../executor/common_kvm_amd64.h' -e 'd' -e '}' | sed -e '/#include \"common_kvm_arm64.h\"/ {' -e 'r ../../executor/common_kvm_arm64.h' -e 'd' -e '}' | sed -e '/#include \"kvm.h\"/ {' -e 'r ../../executor/kvm.h' -e 'd' -e '}' | sed -e '/#include \"kvm.S.h\"/ {' -e 'r ../../executor/kvm.S.h' -e 'd' -e '}' | egrep -v '^[   ]*//' | sed '/^[ 	]*\\/\\/.*/d' | sed 's#[ 	]*//.*##g' >> linux_common.go; echo '`' >> linux_common.go"
//go:generate go fmt linux_common.go

//go:generate bash -c "echo -e '// AUTOGENERATED FROM executor/common_akaros.h\npackage csource\nvar commonHeaderAkaros = `' > akaros_common.go; cat ../../executor/common_akaros.h | sed -e '/#include \"common.h\"/ {' -e 'r ../../executor/common.h' -e 'd' -e '}' | egrep -v '^[   ]*//' | sed '/^[ 	]*\\/\\/.*/d' | sed 's#[ 	]*//.*##g' >> akaros_common.go; echo '`' >> akaros_common.go"
//go:generate go fmt akaros_common.go

//go:generate bash -c "echo -e '// AUTOGENERATED FROM executor/common_bsd.h\npackage csource\nvar commonHeaderFreebsd = `' > freebsd_common.go; cat ../../executor/common_bsd.h | sed -e '/#include \"common.h\"/ {' -e 'r ../../executor/common.h' -e 'd' -e '}' | egrep -v '^[   ]*//' | sed '/^[ 	]*\\/\\/.*/d' | sed 's#[ 	]*//.*##g' >> freebsd_common.go; echo '`' >> freebsd_common.go"
//go:generate go fmt freebsd_common.go

//go:generate bash -c "echo -e '// AUTOGENERATED FROM executor/common_bsd.h\npackage csource\nvar commonHeaderNetbsd = `' > netbsd_common.go; cat ../../executor/common_bsd.h | sed -e '/#include \"common.h\"/ {' -e 'r ../../executor/common.h' -e 'd' -e '}' | egrep -v '^[   ]*//' | sed '/^[ 	]*\\/\\/.*/d' | sed 's#[ 	]*//.*##g' >> netbsd_common.go; echo '`' >> netbsd_common.go"
//go:generate go fmt netbsd_common.go

package csource