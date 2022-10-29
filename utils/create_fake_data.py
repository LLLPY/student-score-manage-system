
from datetime import datetime
from time import time

import os
from random import choices, randint
from secrets import choice


name_set = {'尚', '海', '薛', '屠', '慎', '羊舌', '牟', '路', '皮', '令狐', '缪', '山', '訾', '谷', '鱼', '杭', '梁', '家', '仲', '钟', '幸', '浦', '张', '扶', '梅', '长孙', '庾', '裘', '鲍', '汲', '危', '谯', '荆', '喻', '翟', '韩', '蒙', '笪', '曾', '能', '涂', '洪', '扈', '杨', '殴', '农', '嵇', '狄', '关', '窦', '霍', '石', '相', '叶', '卫', '寿', '邓', '钭', '魏', '松', '昝', '封', '宓', '况', '苍', '计', '陈', '谈', '毋', '颜', '屈', '蓬', '蓟', '逄', '鄂', '车', '姬', '越', '郏', '南宫', '敖', '申屠', '都', '赫连', '暴', '公', '阳', '贲', '井', '湛', '廖', '鞠', '公良', '段干', '佘', '暨', '归', '康', '甄', '晏', '马', '纪', '乔', '夹谷', '羿', '金', '欎', '阚', '荣', '厍', '邵', '柏', '邱', '范', '虞', '弓', '蓝', '琴', '柴', '项', '应', '干', '崔', '章', '詹', '福', '宗政', '耿', '郭', '吴', '于', '竺', '哈', '席', '尉迟', '盛', '艾', '养', '穆', '钮', '富', '万俟', '衡', '佴', '裴', '齐', '东', '季', '辛', '李', '阴', '佟', '傅', '贺', '萧', '後', '步', '凌', '边', '何', '臧', '殳', '麻', '武', '郎', '严', '子车', '蔚', '全', '国', '冉', '易', '雷', '夏', '甘', '楚', '汝', '赵', '段', '红', '慕', '经', '芮', '孔', '祝', '南门', '梁丘', '诸葛', '许', '仇', '邹', '爱', '阎', '谭', '史', '华', '后', '郝', '谢', '吕', '亢', '酆', '柳', '黄', '巫', '盖', '东郭', '言', '郦', '谷梁', '乌', '伯', '房', '王', '管', '文', '班', '亓官', '姚', '阮', '督', '仉', '邬', '帅', '赏', '有', '解', '司空', '宰', '隆', '贝', '景', '包', '淳于', '壤驷', '牧', '满', '钦', '邴', '公冶', '温', '贾', '弘', '闻人', '奚', '明', '龙', '巢', '厉', '尹', '杜', '陆', '胥',
            '利', '戎', '郑', '骆', '倪', '夔', '戴', '诸', '赖', '韦', '余', '申', '闫', '岑', '卜', '巴', '董', '昌', '姜', '巩', '太叔', '邰', '钱', '宋', '韶', '江', '平', '娄', '麴', '宣', '宿', '燕', '乐正', '晋', '徐', '慕容', '岳', '牛', '安', '朱', '程', '公西', '闵', '潘', '双', '秦', '闻', '唐', '侯', '西门', '晁', '束', '汤', '舒', '轩辕', '尤', '龚', '司', '饶', '茅', '端木', '冷', '司寇', '庄', '高', '时', '益', '澹台', '宁', '容', '林', '薄', '俞', '颛孙', '夏侯', '吉', '蔺', '宇文', '逯', '隗', '成', '通', '咸', '游', '毕', '百', '桂', '广', '毛', '充', '桑', '巫马', '翁', '荀', '熊', '漆雕', '费', '鄢', '宦', '桓', '仰', '廉', '钟离', '鲜于', '公羊', '郜', '空', '方', '庞', '宗', '卞', '卓', '姓', '伊', '蒯', '商', '殷', '贡', '符', '和', '焦', '年', '顾', '第五', '籍', '皇甫', '东门', '秋', '司马', '连', '黎', '施', '权', '拓跋', '任', '怀', '沈', '苏', '袁', '池', '孙', '茹', '公孙', '伍', '储', '樊', '别', '沃', '呼延', '乐', '万', '宫', '仲孙', '向', '禹', '田', '蔡', '郁', '匡', '师', '左丘', '惠', '东方', '常', '苗', '白', '司徒', '卢', '莘', '劳', '宰父', '寇', '丁', '曹', '於', '祖', '刘', '查', '胡', '孟', '乜', '左', '周', '柯', '须', '那', '单', '陶', '强', '缑', '百里', '葛', '羊', '糜', '瞿', '习', '终', '阙', '彭', '古', '童', '栾', '璩', '靳', '简', '党', '花', '蒋', '单于', '支', '水', '丰', '冀', '墨', '戈', '微生', '禄', '冯', '融', '上官', '元', '褚', '舄', '凤', '索', '郗', '鲁', '堵', '法', '滕', '印', '云', '米', '濮', '闾丘', '勾', '从', '蒲', '聂', '邢', '莫', '滑', '罗', '汪', '沙', '濮阳', '戚', '伏', '刁', '欧阳', '居', '雍', '祁'}
major_set = {
    '计算机科学',
    '应用数学',
    '信息与计算科学',
    '生物制药',
    '经济管理',
    '哲学',
    '护理学'
}
class_set = {
    '11901',
    '11902',
    '11903',
    '11904',
    '11905',
    '11906',
}
# 创建学生的数据


def create_student_info_data(n=100):
    student_data_path=os.path.join('..','model','data','student.txt')
    score_data_path=os.path.join('..','model','data','score.txt')

    with open(student_data_path, 'w', encoding='utf8') as f_student, open(score_data_path, 'w', encoding='utf8') as f_score:
        student_data = ''
        score_data = ''
        start = 202200001
        for i in range(start, start+n):
            user_type = 1
            num = f'{i}{user_type}'
            name = ''.join(choices(list(name_set), k=randint(2, 3)))
            major = choice(list(major_set))
            class_ = choice(list(class_set))
            birthday = str(datetime.fromtimestamp(
                time()-randint(0, 4*365*24*60*60))).split(' ')[0]
            gender = choice(list(range(2)))
            password = '1234'
            for semester in range(1, 9):
                score = '{},'.format(
                    num)+','.join(map(str, create_score_data()))
                score_data += f'{score},{semester}\n'
                student = f'{num},{name},{major},{class_},{birthday},{gender},{semester},{user_type},{password}\n'
                # print(student)

                student_data += student
        f_student.write(student_data[:-1])  # 去掉最后一个换行
        f_score.write(score_data[:-1])


# 创建分数数据
def create_score_data():
    chinese = randint(50, 100)
    math = randint(60, 100)
    english = randint(30, 100)
    physical = randint(30, 100)
    chemistry = randint(50, 100)
    biology = randint(60, 100)
    sports = randint(80, 100)
    return [chinese, math, english, physical, chemistry, biology, sports]


# 创建老师的数据
def create_teacher_info_data(n=10):
    teacher_data_path=os.path.join('..','model','data','teacher.txt')
    with open(teacher_data_path, 'w', encoding='utf8') as f:
        teacher_data = ''
        user_type = 2
        start = 202200001
        for i in range(start, start+n):
            num = f'{i}{user_type}'
            name = ''.join(choices(list(name_set), k=randint(2, 3)))
            major = choice(list(major_set))
            class_ = choice(list(class_set))
            birthday = str(datetime.fromtimestamp(
                time()-randint(0, 4*365*24*60*60))).split(' ')[0]
            gender = choice(list(range(2)))
            password = '1234'
            teacher = f'{num},{name},{major},{birthday},{gender},{user_type},{class_},{password}\n'
            teacher_data += teacher

        f.write(teacher_data[:-1])


if __name__ == '__main__':
    create_student_info_data()
    create_teacher_info_data()
