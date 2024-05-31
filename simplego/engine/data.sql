drop table pipelines;
drop table components;
create table pipelines ( id INT NOT NULL AUTO_INCREMENT PRIMARY KEY, service varchar(255) NOT NULL, hook varchar(255) NOT NULL, type varchar(255) NOT NULL, app_id int NOT NULL, platform int NOT NULL, override_id int, extra LONGTEXT, content LONGTEXT, version TEXT, create_time TIMESTAMP DEFAULT 0, update_time TIMESTAMP DEFAULT now() ON UPDATE now() ) ENGINE = InnoDB;
create table components (id INT NOT NULL AUTO_INCREMENT PRIMARY KEY, service varchar(255) NOT NULL, hook varchar(255) NOT NULL, name varchar(255) NOT NULL, version TEXT, type varchar(255) NOT NULL, parameter_type TEXT, content LONGTEXT, create_time TIMESTAMP DEFAULT now(), update_time TIMESTAMP DEFAULT now() on UPDATE now()) ENGINE=InnoDB;

INSERT INTO pipelines (service, hook, type, app_id, platform, content, version) VALUES('service1', 'grouping', 'code', 1128, 1, 'WrappedC2()+WrappedC1(bound=50)', 'online');

INSERT INTO pipelines (service, hook, type, app_id, platform, content, version) VALUES('service1', 'grouping', 'code', 1128, 4, 'WrappedAny(bound=50,a=30,c="config")+WrappedC2()', 'online');

INSERT INTO pipelines (service, hook, type, app_id, platform, content, version) VALUES('service2', 'filter', 'code', 1128, 1, 'ISPMatch(isp=1)+ProvSame()', 'online');
INSERT INTO pipelines (service, hook, type, app_id, platform, content, version) VALUES('service2', 'filter', 'code', 2329, 1, 'ISPMatch(isp=1)+ProvNotSame()', 'online');


INSERT INTO components (service, hook, name, type, content) VALUES(
    'service1', 
    'grouping', 
    'WrappedAny',
    'code',
    'import (
		"fmt"
	)

	func WrappedAny(params ...interface{}) func(any) (int, int, int, int) {
		x := params[0].(int)
		fmt.Printf("WrappedAny params: %v\\n", params)
		return func(s any) (id, reset, code, quit int) {
            fmt.Println("host_arch:", s.(*Mixed).HostArch)
			if s.(*Mixed).Concurrent > x {
				return 2003, 1, 233, 1
			}
			if s.(*Mixed).Platform == 4 {
				return 2002, 0, 234, 1
			}
			return 2002, 0, 235, 0
		}
	}
    '
);


INSERT INTO components (service, hook, name, type, content) VALUES(
    'service1', 
    'grouping', 
    'WrappedC2',
    'code',
    '
	func WrappedC2() func(any) (int, int, int, int) {
		return func(s any) (id, reset, code, quit int) {
			if s.(*Mixed).Concurrent == 137 {
				return 2003, 1, 300, 1
			}
			return 2002, 0, 235, 0
		}
	}
    '
);

INSERT INTO components (service, hook, name, type, content) VALUES(
    'service2', 
    'filter', 
    'ISPMatch',
    'code',
    '
	import (
		"fmt"
	)

	func ISPMatch(params ...interface{}) func(any) (int, int, int, int) {
		x := params[0].(int)
		return func(s any) (id, reset, code, quit int) {
			if s.(*PMPeerInfo).ISP != x {
				return 0, 0, 0, 1
			}
			return 1, 0, 0, 0
		}
	}
    '
);

INSERT INTO components (service, hook, name, type, content) VALUES(
    'service2', 
    'filter', 
    'ProvSame',
    'code',
    '
	import (
		"fmt"
	)

	func ProvSame(params ...interface{}) func(any) (int, int, int, int) {
		return func(s any) (id, reset, code, quit int) {
			p := s.(*PMPeerInfo)
			if p.RequesterProv != p.Prov {
				return 0, 0, 0, 1
			}
			return 1, 0, 0, 0
		}
	}
    '
);

INSERT INTO components (service, hook, name, type, content) VALUES(
    'service2', 
    'filter', 
    'ProvNotSame',
    'code',
    '
	import (
		"fmt"
	)

	func ProvNotSame(params ...interface{}) func(any) (int, int, int, int) {
		return func(s any) (id, reset, code, quit int) {
			p := s.(*PMPeerInfo)
			if p.RequesterProv == p.Prov {
				return 0, 0, 0, 1
			}
			return 1, 0, 0, 0
		}
	}
    '
);

INSERT INTO components (service, hook, name, type, content) VALUES(
    'service1', 
    'grouping', 
    'WrappedC1',
    'code',
    '
	import (
		"fmt"
		"encoding/json"
	)

	type Version struct {
		Name string
	}

	func WrappedC1(bound int) func(any) (int, int, int, int) {
		x := bound
		return func(s any) (id, reset, code, quit int) {
			if s.(*Mixed).Concurrent > x {
				return 2003, 1, 233, 1
			}
			if s.(*Mixed).Platform == 4 {
				return 2002, 0, 234, 1
			}
			return 2002, 0, 235, 0
		}
	}
    '
);

