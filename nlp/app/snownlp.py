# text = u'''
# 自然语言处理是计算机科学领域与人工智能领域中的一个重要方向。
# 它研究能实现人与计算机之间用自然语言进行有效通信的各种理论和方法。
# 自然语言处理是一门融语言学、计算机科学、数学于一体的科学。
# 因此，这一领域的研究将涉及自然语言，即人们日常使用的语言，
# 所以它与语言学的研究有着密切的联系，但又有重要的区别。
# 自然语言处理并不是一般地研究自然语言，
# 而在于研制能有效地实现自然语言通信的计算机系统，
# 特别是其中的软件系统。因而它是计算机科学的一部分。
# '''
#
# s = SnowNLP(text)
#
# s.keywords(3)	# [u'语言', u'自然', u'计算机']
#
# s.summary(3)	# [u'因而它是计算机科学的一部分',
# #  u'自然语言处理是一门融语言学、计算机科学、
# #	 数学于一体的科学',
# #  u'自然语言处理是计算机科学领域与人工智能
# #	 领域中的一个重要方向']

from flask import request, jsonify
from snownlp import SnowNLP

from . import bp


@bp.route('/analysis', methods=['POST'])
def analysis():
    try:
        data = request.get_json()
        if data is None:
            return jsonify({}), 400
        text = data['text']
        s = SnowNLP(text)
        keyword = s.keywords(5)
        summary = s.summary(5)
        return jsonify({"keyword": keyword, "summary": summary}), 200
    except Exception as e:
        return jsonify({"error": e}), 500
