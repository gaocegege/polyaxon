# Generated by Django 2.1.3 on 2018-11-26 12:41

from django.db import migrations, models
import django.db.models.deletion


def migrate_iteration_experiments(apps, schema_editor):
    ExperimentGroupIteration = apps.get_model('db', 'ExperimentGroupIteration')
    Experiment = apps.get_model('db', 'Experiment')

    for iteration in ExperimentGroupIteration.objects.all():
        if iteration.data and 'experiment_ids' in iteration.data:
            experiment_ids = Experiment.objects.filter(
                id__in=iteration.data['experiment_ids']).values_list('id', flat=True)
            iteration.experiments.set(experiment_ids)


class Migration(migrations.Migration):

    dependencies = [
        ('db', '0014_add_owner_and_token'),
    ]

    operations = [
        migrations.AddField(
            model_name='experimentgroupiteration',
            name='experiments',
            field=models.ManyToManyField(blank=True, related_name='_experimentgroupiteration_experiments_+', to='db.Experiment'),
        ),
        migrations.AlterField(
            model_name='owner',
            name='content_type',
            field=models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, related_name='+', to='contenttypes.ContentType'),
        ),
        migrations.AlterField(
            model_name='project',
            name='owner',
            field=models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, related_name='projects', to='db.Owner'),
        ),
        migrations.RunPython(migrate_iteration_experiments),
    ]