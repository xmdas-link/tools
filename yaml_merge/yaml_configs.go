package yaml_merge

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ghodss/yaml"
	"io/ioutil"
	"os"
	"strings"
)

func New() *YamlConfigs {
	return &YamlConfigs{}
}

type YamlConfigs struct {
	data map[string]interface{}
}

// 添加yaml文件
func (y *YamlConfigs) AddFile(file string) error {

	// 读取文件
	f, err := os.Open(file)
	if err != nil {
		return err
	}

	out, errRead := ioutil.ReadAll(f)
	if errRead != nil {
		return errRead
	}

	// 转json
	jsonBytes, jsonErr := yaml.YAMLToJSON(out)
	if jsonErr != nil {
		return jsonErr
	}

	// 将JSON转map
	data := map[string]interface{}{}
	if err := json.Unmarshal(jsonBytes, &data); err != nil {
		return err
	}

	return y.Merge(data)
}

func (y *YamlConfigs) Merge(data map[string]interface{}) error {

	if y.data == nil {
		y.data = data
		return nil
	}

	for key, d := range data {
		if mergeData, err := MergeData([]string{key}, y.data[key], d); err != nil {
			return err
		} else {
			y.data[key] = mergeData
		}
	}

	return nil
}

func MergeData(paths []string, org interface{}, new interface{}) (interface{}, error) {

	if org == nil {
		return new, nil
	}

	// 检查数据类型
	switch org.(type) {
	case map[string]interface{}:
		if orgMap, ok := org.(map[string]interface{}); ok {
			if newMap, ok2 := new.(map[string]interface{}); ok2 {
				for key, d := range newMap {
					if mergeData, err := MergeData(append(paths, key), orgMap[key], d); err != nil {
						return nil, err
					} else {
						orgMap[key] = mergeData
					}
				}
				return orgMap, nil
			} else {
				return nil, errors.New(fmt.Sprintf("要覆盖的参数%s的数值不匹配，应该是map[string]interface{}, 实际不是", strings.Join(paths, ".")))
			}
		} else {
			return nil, errors.New(fmt.Sprintf("参数%s被判断为map[string]interface{}，但是数据转换失败", strings.Join(paths, ".")))
		}
	default:
		return new, nil
	}

}

func (y *YamlConfigs) SaveAs(file string) error {

	if y.data == nil {
		return errors.New("没有可以输出的内容")
	}

	// map转JSON
	jsonBytes, jsonErr := json.Marshal(y.data)
	if jsonErr != nil {
		return jsonErr
	}

	// JSON转YAML
	data, yamlErr := yaml.JSONToYAML(jsonBytes)
	if yamlErr != nil {
		return yamlErr
	}

	// 覆盖写入
	yamlFile, errFile := os.OpenFile(file, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)
	if errFile != nil {
		return errFile
	}

	yamlFile.Write(data)

	return nil
}

func (y *YamlConfigs) YamlToByte() ([]byte, error) {
	var data []byte

	if y.data == nil {
		return data, errors.New("没有可以输出的内容")
	}

	// map转JSON
	jsonBytes, jsonErr := json.Marshal(y.data)
	if jsonErr != nil {
		return data, jsonErr
	}

	// JSON转YAML
	data, yamlErr := yaml.JSONToYAML(jsonBytes)
	if yamlErr != nil {
		return data, yamlErr
	}
	return data, nil
}
